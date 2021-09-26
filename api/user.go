package api

import (
	"context"
	"net/http"

	"curd_demo/model"
	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

func UserCreate(ctx *gin.Context) {
	req := new(model.UserCreateRequest)
	err := ctx.BindJSON(req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}
	resp, err := NewUser(db).Create(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}
	ctx.JSON(http.StatusOK, resp)
}

type User interface {
	Create(ctx context.Context, req *model.UserCreateRequest) (*model.UserCreateResponse, error)
}

func NewUser(db *gorm.DB) User {
	return &UserEntry{db: db}
}

type UserEntry struct {
	db *gorm.DB
}

func (entry *UserEntry) Create(ctx context.Context, req *model.UserCreateRequest) (*model.UserCreateResponse, error) {
	user := model.NewUser(req)
	err := entry.db.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return model.NewUserCreateResponse(user), nil
}
