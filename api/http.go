package api

import (
	"curd_demo/config"
	"github.com/gin-gonic/gin"
)

func StartHttpService() {
	engine := gin.Default()
	SetUpRoutes(engine)
	err := engine.Run(config.Hub.HttpSetting.Addr)
	if err != nil {
		panic(err)
	}
}

func SetUpRoutes(engine *gin.Engine) {
	engine.GET("/hello", Hello)
	engine.POST("/user/create", UserCreate)
}
