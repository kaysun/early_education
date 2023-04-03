package content

import (
	"github.com/circle/early_education/transport/content/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Homepage 首页频道
func Homepage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.HomepageResp{},
	})
}

// ModContents 获取模块内容列表
func ModContents(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.ModContentsResp{},
	})
}
