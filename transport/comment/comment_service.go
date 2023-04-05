package comment

import (
	"github.com/circle/early_education/transport/comment/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// FollowAndPlayContent 跟玩内容
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
