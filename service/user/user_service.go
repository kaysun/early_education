package user

import (
	"github.com/circle/early_education/infra"
	"github.com/circle/early_education/model"
)

// UserService 用户服务
type UserService struct {
	userDao infra.UserDAO
}

func NewUserService(userDao infra.UserDAO) UserService {
	return UserService{
		userDao: userDao,
	}
}

// Register 注册
func (u UserService) Register(user model.User) (int, error) {
	return u.userDao.AddUser(user)
}

// Login 登录
func (u UserService) Login(user model.User) (int, error) {
	userID, err := u.userDao.GetUserID(user)
	if err != nil {
		return 0, err
	}

	return userID, nil
}
