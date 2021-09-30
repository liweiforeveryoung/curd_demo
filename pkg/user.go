package pkg

import (
	"context"

	"curd_demo/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

//go:generate mockgen -destination user_mock.go -package pkg -source user.go User
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
