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

// AgeStage 获取用户的年龄阶段
func (w User) AgeStage() AgeStage {
	// duration是纳秒
	duration := time.Since(w.Birthday)
	month := duration / 1000 / 1000 / 1000 / 3600 / 24 / 30
	year := month / 12
	mod := month % 12
	if year > 6 {
		return YearMoreThan6
	} else if year >= 5 {
		return Year5To6
	} else if year >= 4 {
		return Year4To5
	} else if year >= 3 {
		return Year3To4
	} else if year >= 2 {
		return Year2To3
	} else if year >= 1 {
		return Year1To2
	} else if mod >= 10 && mod <= 12 {
		return Month10To12
	} else if mod >= 7 && mod <= 9 {
		return Month7To9
	} else if mod >= 0 && mod <= 6 {
		return Month0To6
	}
	return 0
}
