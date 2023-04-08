package course

import (
	"github.com/circle/early_education/transport/course/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// TeacherList 教师列表
func TeacherList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.TeacherListResp{},
	})
}
