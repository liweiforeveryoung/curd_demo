package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func Hello(context *gin.Context) {
	context.String(http.StatusOK, "hello world")
}
