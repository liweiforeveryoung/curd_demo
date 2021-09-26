package apis

import (
	"net/http"

	"curd_demo/model"
	"github.com/gin-gonic/gin"
)

func UserCreate(ctx *gin.Context) {
	userCreateRequest := new(model.UserCreateRequest)
	err := ctx.BindJSON(userCreateRequest)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	user := model.NewUser(userCreateRequest)
	err = db.WithContext(ctx).Create(user).Error
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, model.NewUserCreateResponse(user))
}
