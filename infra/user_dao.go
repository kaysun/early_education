package infra

import (
	"github.com/circle/early_education/model"
	"gorm.io/gorm"
)

type UserDAO struct {
	// orm 对象数据映射关系
	MysqlProxy *gorm.DB
}

// NewUserDAO 创建用户数据访问层
func NewUserDAO() UserDAO {
	return UserDAO{
		MysqlProxy: &gorm.DB{
			Config:       nil,
			Error:        nil,
			RowsAffected: 0,
			Statement:    nil,
		},
	}
}

func (u UserDAO) GetUserID(user model.User) (int, error) {
	// 通过MysqlProxy 操作数据库
	var newUser model.User
	err := u.MysqlProxy.
		Select("user_id").
		// 防止sql注入
		Where("name=? and password=?", user.Name, user.Password).
		First(&newUser).
		Error
	// 如果能取出来，userid一定大于0，说明有这个用户
	return int(newUser.UserID), err
}

func (u UserDAO) AddUser(user model.User) (int, error) {
	err := u.MysqlProxy.
		Create(&user).
		Error
	return int(user.UserID), err
}
