package pkg

import (
	"context"
	"fmt"
	"testing"

	"curd_demo/config"
	"curd_demo/model"
	"curd_demo/util"
	gomysql "github.com/go-sql-driver/mysql"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func initDB() *gorm.DB {
	config.Initialize()
	dsnCfg, err := gomysql.ParseDSN(config.Hub.DBSetting.MysqlDSN)
	if err != nil {
		panic(err)
	}
	dbName := dsnCfg.DBName
	// 在 dsn 中不指定 dbname
	dsnCfg.DBName = ""

	db, err := gorm.Open(mysql.Open(dsnCfg.FormatDSN()), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	// 开启 debug 模式 方便看到每次执行时的 sql 语句
	db.Logger = db.Logger.LogMode(logger.Info)
	drop := fmt.Sprintf("DROP DATABASE IF EXISTS %s;", dbName)
	create := fmt.Sprintf("CREATE DATABASE %s DEFAULT CHARSET = utf8mb4 COLLATE = utf8mb4_unicode_ci;", dbName)
	use := fmt.Sprintf(`USE %s;`, dbName)
	migrations := util.FolderContentLoad(config.ProjectName, config.MigrationsFolderName)
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
	return db
}

func TestUserEntry_Create(t *testing.T) {
	db := initDB()
	entry := UserEntry{db: db}
	user := &model.User{
		Name: "levy",
		Age:  18,
		Sex:  model.MAN,
	}
	userCreateReq := model.NewUserCreateRequest(user)
	userCreateResp, err := entry.Create(context.TODO(), userCreateReq)
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
