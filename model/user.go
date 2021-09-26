package model

import (
	"math/rand"

	"github.com/jinzhu/copier"
)

type Sex int8

const (
	UNKNOWN = 0
	MAN     = 1
	WOMAN   = 2
)

type User struct {
	Id int64 `gorm:"column:id" json:"-"`
	// 用户 id
	UserId int64 `gorm:"column:user_id" json:"user_id"`
	// 用户昵称
	Name string `gorm:"column:name" json:"name"`
	// 用户年龄
	Age int32 `gorm:"column:age" json:"age"`
	// 用户性别
	Sex       Sex   `gorm:"column:sex" json:"sex"`
	DeletedAt int64 `gorm:"column:deleted_at" json:"-"`
	CreatedAt int64 `gorm:"column:created_at;autoCreateTime" json:"-"`
	UpdatedAt int64 `gorm:"column:updated_at;autoUpdateTime" json:"-"`
}

// DeepCopy 对 User 进行深度拷贝
func (m *User) DeepCopy() *User {
	cp := new(User)
	_ = DeepCopy(cp, m)
	return cp
}

func (m *User) TableName() string {
	return "users"
}

func NewUser(req *UserCreateRequest) *User {
	return &User{
		UserId: rand.Int63(), // 随便写的
		Name:   req.User.Name,
		Age:    req.User.Age,
		Sex:    req.User.Sex,
	}
}

func NewUserCreateResponse(user *User) *UserCreateResponse {
	return &UserCreateResponse{User: user}
}

func NewUserCreateRequest(user *User) *UserCreateRequest {
	return &UserCreateRequest{User: user}
}

func DeepCopy(toValue interface{}, fromValue interface{}) error {
	return copier.CopyWithOption(toValue, fromValue, copier.Option{
		IgnoreEmpty: true,
		DeepCopy:    true,
	})
}

type UserCreateRequest struct {
	User *User `json:"user"`
}

type UserCreateResponse struct {
	User *User `json:"user"`
}
