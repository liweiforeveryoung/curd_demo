package api

import (
	"github.com/liweiforeveryoung/curd_demo/dep"
	"github.com/liweiforeveryoung/curd_demo/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserCreate(ctx *gin.Context) {
	req := new(model.UserCreateRequest)
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := dep.Hub.User.Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, resp)
}
