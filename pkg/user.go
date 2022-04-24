package pkg

import (
	"context"
	"github.com/liweiforeveryoung/curd_demo/model"

	"github.com/pkg/errors"
	"go.uber.org/dig"
	"gorm.io/gorm"
)

//go:generate mockgen -destination user_mock.go -package pkg -source user.go User
type User interface {
	Create(ctx context.Context, req *model.UserCreateRequest) (*model.UserCreateResponse, error)
}

func NewUser(entry UserEntry) User {
	return &entry
}

type UserEntry struct {
	dig.In
	DB *gorm.DB
}

func (entry *UserEntry) Create(ctx context.Context, req *model.UserCreateRequest) (*model.UserCreateResponse, error) {
	user := model.NewUser(req)
	err := entry.DB.WithContext(ctx).Create(user).Error
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return model.NewUserCreateResponse(user), nil
}
