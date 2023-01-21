package routers

import (
	"github.com/circle/early_education/transport/user"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/user")
	{
		// 注册
		userRouter.POST("/register", user.Resister)
		// 登录
		userRouter.POST("/login", user.Login)
	}
	return r
}
