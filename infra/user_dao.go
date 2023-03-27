package infra

import (
	"github.com/circle/early_education/global"
	"github.com/circle/early_education/model"
)

type UserDAO struct {
}

func (u UserDAO) GetUserID(user model.User) (int, error) {
	// 通过MysqlProxy 操作数据库
	var newUser model.User
	err := global.MysqlProxy.
		Select("user_id").
		// 防止sql注入
		Where("name=? and password=?", user.Name, user.Password).
		First(&newUser).
		Error
	// 如果能取出来，userid一定大于0，说明有这个用户
	return int(newUser.UserID), err
}

func (u UserDAO) AddUser(user model.User) (int, error) {
	err := global.MysqlProxy.
		Create(&user).
		Error
	return int(user.UserID), err
}
