package main

import (
	"math/rand"
	"net/http"
	"time"

	"curd_demo/model"
	"github.com/gin-gonic/gin"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func main() {
	var err error
	dsn := "root:@tcp(127.0.0.1:3306)/curd_db?charset=utf8mb4&parseTime=True&loc=Local"
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	rand.Seed(time.Now().Unix())
	engine := gin.Default()
	SetUpRoutes(engine)
	err = engine.Run(":3344")
	if err != nil {
		panic(err)
	}
}

func SetUpRoutes(engine *gin.Engine) {
	engine.GET("/hello", Hello)
	engine.POST("/user/create", UserCreate)
}

func Hello(context *gin.Context) {
	context.String(http.StatusOK, "hello world")
}

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
