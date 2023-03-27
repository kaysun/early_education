package model

import "time"

// User 用户
type User struct {
	// UserID 用户id
	UserID uint64 `gorm:"column:user_id"`
	// Name 用户名，全局唯一
	Name string `gorm:"column:name"`
	// Password 密码，存储密文
	Password string `gorm:"column:password"`
	// Nick 昵称
	Nick string `gorm:"column:nick"`
	// Birthday 出生日期，用于判定
	Birthday time.Time `gorm:"column:birthday"`
	// Phone 手机号
	Phone int `gorm:"column:phone"`
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time"`
	// UpdateTime 修改时间
	UpdateTime time.Time `gorm:"column:update_time"`
}

// TableName 表名
func (w User) TableName() string {
	return "user"
}
