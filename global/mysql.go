package global

import "gorm.io/gorm"

// MysqlProxy MySQL代理
var MysqlProxy *gorm.DB

func InitMySQL() {
	// TODO 连接mysql
}
