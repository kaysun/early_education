package course

import (
	"github.com/circle/early_education/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

// All 所有的课程列表
func All(context *gin.Context) {
	//timeVal1, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-22 09:30:00", time.Local)
	//timeVal2, _ := time.ParseInLocation("2006-01-02 15:04:05", "2023-01-22 11:00:00", time.Local)

	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": []model.Course{},
	})
}
