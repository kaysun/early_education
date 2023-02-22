package model

// User 用户
type User struct {
	UserID uint64 `gorm:"column:user_id"`
	// Name 用户名，全局唯一
	Name string `gorm:"column:name"`
	// Password 密码，存储密文
	Password string `gorm:"column:password"`
	// Birthday 出生日期，用于判定
	Birthday string `gorm:"column:birthday"`
}

// TableName 表名
func (w User) TableName() string {
	return "user"
}
