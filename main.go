package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	SetUpRoutes(engine)
	err := engine.Run(":3344")
	if err != nil {
		panic(err)
	}
}

func SetUpRoutes(engine *gin.Engine) {
	engine.GET("/hello", Hello)
}

func Hello(context *gin.Context) {
	context.String(http.StatusOK, "hello world")
}
