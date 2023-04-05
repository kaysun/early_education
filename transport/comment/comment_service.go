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
