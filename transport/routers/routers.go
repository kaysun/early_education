package routers

import (
	"github.com/circle/early_education/transport/comment"
	"github.com/circle/early_education/transport/content"
	"github.com/circle/early_education/transport/course"
	"github.com/circle/early_education/transport/user"
	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {
	r := gin.Default()

	// 用户
	userRouter := r.Group("/user")
	{
		// 注册 /user/register
		userRouter.POST("/register", user.Resister)
		// 登录 /user/login
		userRouter.POST("/login", user.Login)
	}

	// 课程
	courseRouter := r.Group("/course")
	{
		// 课程列表，默认返回7天的 /course/all
		courseRouter.GET("/all", course.All)
		// 获取可约课程列表 /course/reservable/list
		courseRouter.GET("/reservable/list", course.ReservableList)
		// 约课 /course/booking
		courseRouter.POST("/booking", course.Booking)
		// 取消约课 /course/cancel/booking
		courseRouter.POST("/cancel/booking", course.CancelBooking)
	}

	// 内容
	contentRouter := r.Group("/content")
	{
		// 首页列表 /content/homepage
		contentRouter.GET("/homepage", content.Homepage)
		// 分页获取模块下的所有内容 /content/mod/contents
		contentRouter.GET("/mod/contents", content.ModContents)
		// 浏览内容 /content/view
		contentRouter.POST("/view", content.ViewContent)
		// 收藏内容 /content/collect
		contentRouter.POST("/collect", content.CollectContent)
		// 取消收藏内容 /content/cancel/collect
		contentRouter.POST("/cancel/collect", content.CancelCollectContent)
		// 点赞内容 /content/up
		contentRouter.POST("/up", content.UpContent)
		// 取消点赞内容 /content/cancel/up
		contentRouter.POST("/cancel/up", content.CancelUpContent)
		// 跟玩内容 /content/follow_and_play
		contentRouter.POST("/follow_and_play", comment.FollowAndPlayContent)
		// 官方提供+推荐+自己创建的+用户参与的育儿专题列表
		contentRouter.GET("/topics", content.Topics)
		// 创建专题
		contentRouter.POST("/topic/create", content.CreateTopic)
	}

	// 评论
	commentRouter := r.Group("/comment")
	{
		// 分页获取内容下的评论列表 /comment/list
		commentRouter.GET("/list", comment.List)
		// 发表评论 /comment/create
		commentRouter.POST("/create", comment.Create)
	}

	return r
}
