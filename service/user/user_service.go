package user

import (
	"github.com/circle/early_education/infra"
	"github.com/circle/early_education/model"
)

// UserService 用户服务
type UserService struct {
}

// Register 注册
func (UserService) Register() error {
	return nil
}

// Login 登录
func (u UserService) Login(user model.User) (bool, error) {
	userDAO := infra.NewUserDAO()
	userID, err := userDAO.GetUserID(user)
	if err != nil {
		return false, err
	}

	return userID > 0, nil
}
