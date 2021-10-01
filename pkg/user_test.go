package pkg

import (
	"context"

	"curd_demo/model"
)

func (s *SuiteTest) TestUserEntry_Create() {
	user := &model.User{
		Name: "levy",
		Age:  18,
		Sex:  model.MAN,
	}
	userCreateReq := model.NewUserCreateRequest(user)
	userCreateResp, err := s.userEntry.Create(context.TODO(), userCreateReq)
	s.NoError(err)
	s.Equal(user.Name, userCreateResp.User.Name)
	s.Equal(user.Age, userCreateResp.User.Age)
	s.Equal(user.Sex, userCreateResp.User.Sex)
	// 检查数据库
	userFromDB := user.DeepCopy()
	err = s.db.First(userFromDB, user).Error
	s.NoError(err)
	s.NotEqual(int64(0), userFromDB.UserId)
}
