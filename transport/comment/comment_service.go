package comment

import (
	"github.com/circle/early_education/transport/comment/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FollowAndPlayContent 跟玩内容。注意，直接调用跟创建评论一样的service即可。
func FollowAndPlayContent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.CommentResp{},
	})
}

// List 分页获取内容下的评论列表
func List(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.CommentListResp{},
	})
}

// Create 发表评论
func Create(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.CommentResp{},
	})
}

// Reply 回复评论
func Reply(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.CommentResp{},
	})
}

// Up 点赞评论
func Up(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": struct{}{},
	})
}

// CancelUp 取消点赞评论
func CancelUp(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": struct{}{},
	})
}
