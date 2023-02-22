package routers

import (
	"github.com/circle/early_education/transport/course"
	"github.com/circle/early_education/transport/user"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()
	userRouter := r.Group("/user")
	{
		// 注册 /user/register?
		userRouter.POST("/register", user.Resister)
		// 登录 /user/login?
		userRouter.POST("/login", user.Login)
	}

	courseRouter := r.Group("/course")
	{
		// 课程列表，默认返回7天的 /course/list?pageNo=1&size=10
		courseRouter.GET("/list", course.List)
	}

	return r
}
