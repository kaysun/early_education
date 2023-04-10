package global

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// MysqlProxy MySQL代理
var MysqlProxy *gorm.DB

func InitMySQL() {
	dsn := "user:pass@tcp(127.0.0.1:3306)/education?charset=utf8mb4&parseTime=True&loc=Local"
	var err error
	MysqlProxy, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println(fmt.Sprintf("初始化mysql失败，err=%+v", err))
	}
}
