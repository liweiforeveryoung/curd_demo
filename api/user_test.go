package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"curd_demo/dep"
	"curd_demo/model"
	"curd_demo/pkg"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestUserCreate(t *testing.T) {
	g := gin.Default()
	SetUpRoutes(g)

	user := &model.User{
		Name: "levy",
		Age:  18,
		Sex:  model.MAN,
	}
	userCreateReq := model.NewUserCreateRequest(user)
	contentBytes, _ := json.Marshal(userCreateReq)
	reader := bytes.NewReader(contentBytes)

	req := httptest.NewRequest("POST", "/user/create", reader)
	w := httptest.NewRecorder()

	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userMock := pkg.NewMockUser(ctrl)
	dep.Hub.User = userMock
	// 可以用 gomock.Any() 代表任何类型的参数
	userMock.EXPECT().Create(gomock.Any(), userCreateReq).Return(new(model.UserCreateResponse), nil)

	g.ServeHTTP(w, req)

	// 只简单检查一下 StatusCode
	assert.Equal(t, http.StatusOK, w.Code)
}
