package api

import (
	"curd_demo/config"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
)

func StartHttpService() error {
	logrus.Info("start http service ...")
	engine := gin.Default()
	SetUpRoutes(engine)
	err := engine.Run(config.Hub.HttpSetting.Addr)
	if err != nil {
		return errors.WithStack(err)
	}
	return nil
}

func SetUpRoutes(engine *gin.Engine) {
	engine.GET("/hello", Hello)
	engine.POST("/user/create", UserCreate)
}
