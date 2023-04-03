package content

import (
	"github.com/circle/early_education/transport/course/proto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// Homepage 首页频道
func Homepage(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.CancelBookingCourseResp{BookingID: 0},
	})
}
