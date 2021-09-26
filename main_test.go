package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"curd_demo/model"
	"curd_demo/util"
	"github.com/gin-gonic/gin"
	gomysql "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func TestHello(t *testing.T) {
	g := gin.Default()
	SetUpRoutes(g)

	req := httptest.NewRequest("GET", "/hello", nil)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}

const ProjectName = "curd_demo"

func initDB() {
	// 当 clientFoundRows 为 true, db.Update().RowsAffected 返回的是匹配到的记录行数
	// 当 clientFoundRows 为 false, db.Update().RowsAffected 返回的是实际更新的记录行数
	dsn := "root:@tcp(127.0.0.1:3306)/curd_db_test?charset=utf8mb4&parseTime=True&loc=Local&clientFoundRows=true"

	dsnCfg, err := gomysql.ParseDSN(dsn)
	if err != nil {
		panic(err)
	}
	dbName := dsnCfg.DBName
	// 在 dsn 中不指定 dbname
	dsnCfg.DBName = ""

	db, err = gorm.Open(mysql.Open(dsnCfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 开启 debug 模式 方便看到每次执行时的 sql 语句
	db.Logger = db.Logger.LogMode(logger.Info)
	drop := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)
	create := fmt.Sprintf("CREATE DATABASE %s DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;", dbName)
	use := fmt.Sprintf(`USE %s;`, dbName)
	migrations := LoadMysqlPath("migrations")
	err = db.Exec(drop).Exec(create).Exec(use).Error
	if err != nil {
		panic(err)
	}
	for _, migrate := range migrations {
		err = db.Exec(migrate).Error
		if err != nil {
			panic(err)
		}
	}
}

func TestUserCreate(t *testing.T) {
	initDB()

	g := gin.Default()
	SetUpRoutes(g)

	user := &model.User{
		Name: "levy",
		Age:  18,
		Sex:  model.MAN,
	}
	userCreateReq := model.NewUserCreateRequest(user)
	contentBytes, _ := json.Marshal(userCreateReq)
	reader := bytes.NewReader(contentBytes)

	req := httptest.NewRequest("POST", "/user/create", reader)
	w := httptest.NewRecorder()
	g.ServeHTTP(w, req)

	// 检查 resp
	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode)
	userCreateResp := new(model.UserCreateResponse)
	err := BindResp(resp, userCreateResp)
	assert.NoError(t, err)
	assert.Equal(t, user.Name, userCreateResp.User.Name)
	assert.Equal(t, user.Age, userCreateResp.User.Age)
	assert.Equal(t, user.Sex, userCreateResp.User.Sex)
	// 检查数据库
	userFromDB := user.DeepCopy()
	err = db.First(userFromDB, user).Error
	assert.NoError(t, err)
	assert.NotEqual(t, int64(0), userFromDB.UserId)
}

func BindResp(resp *http.Response, obj interface{}) error {
	defer resp.Body.Close()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("ReadAll err[%w]", err)
	}
	err = json.Unmarshal(content, obj)
	if err != nil {
		return fmt.Errorf("unmarshal err[%w]", err)
	}
	return nil
}

// LoadMysqlPath 将 dirName 每个文件中的 content 以 string 的形式 load 出来
func LoadMysqlPath(dirName string) []string {
	dirPath, err := util.DirOrFilePathFromProject(ProjectName, dirName)
	if err != nil {
		panic(err)
	}
	fileNames, err := util.FileNamesInDir(dirPath)
	if err != nil {
		panic(err)
	}
	return LoadMysqlFiles(fileNames)
}

func LoadMysqlFiles(sqlFileNames []string) []string {
	contents := make([]string, 0, len(sqlFileNames))
	for _, name := range sqlFileNames {
		contents = append(contents, LoadMysqlFile(name))
	}
	return contents
}

func LoadMysqlFile(sqlFileName string) string {
	sqlFilePath, err := util.DirOrFilePathFromProject(ProjectName, sqlFileName)
	if err != nil {
		panic(err)
	}
	content, err := os.ReadFile(sqlFilePath)
	if err != nil {
		panic(err)
	}
	return string(content)
}
