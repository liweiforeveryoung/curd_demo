package api

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"curd_demo/model"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
)

func (s *SuiteTest) TestUserCreate() {
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
	// 可以用 gomock.Any() 代表任何类型的参数
	s.userMock.EXPECT().Create(gomock.Any(), userCreateReq).Return(new(model.UserCreateResponse), nil)

	g.ServeHTTP(w, req)
	// 只简单检查一下 StatusCode
	s.Equal(http.StatusOK, w.Code)
}
