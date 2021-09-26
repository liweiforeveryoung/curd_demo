package apis

import (
	"curd_demo/config"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartHttpService() {
	var err error
	db, err = gorm.Open(mysql.Open(config.Hub.DBSetting.MysqlDSN), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	engine := gin.Default()
	SetUpRoutes(engine)
	err = engine.Run(config.Hub.HttpSetting.Addr)
	if err != nil {
		panic(err)
	}
}

func SetUpRoutes(engine *gin.Engine) {
	engine.GET("/hello", Hello)
	engine.POST("/user/create", UserCreate)
}
