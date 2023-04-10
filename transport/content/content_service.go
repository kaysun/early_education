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

// Topics 官方提供+推荐+自己创建的育儿专题列表
func Topics(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.TopicsResp{},
	})
}

// CreateTopic 创建专题
func CreateTopic(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.TopicCreateResp{},
	})
}

// RecommendTopic 推荐专题
func RecommendTopic(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.TopicRecommendResp{},
	})
}

// ViewContent 浏览内容
func ViewContent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.ContentOperationResp{},
	})
}

// CollectContent 收藏内容
func CollectContent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.ContentOperationResp{},
	})
}

// CancelCollectContent 取消收藏内容
func CancelCollectContent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.ContentOperationResp{},
	})
}

// UpContent 点赞内容
func UpContent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.ContentOperationResp{},
	})
}

// CancelUpContent 取消点赞内容
func CancelUpContent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.ContentOperationResp{},
	})
}

// UploadContent 上传内容
func UploadContent(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.ContentUploadResp{},
	})
}
