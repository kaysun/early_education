package course

import (
	"github.com/circle/early_education/transport/course/proto"
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
		"data": proto.GetAllCoursesResp{
			Courses: []proto.Course{},
		},
	})
}

// ReservableList 可约课程列表
func ReservableList(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.ReservableListResp{
			ReservableCourses: []proto.ReservableCourse{},
		},
	})
}

// Booking 约课
func Booking(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.BookingCourseResp{BookingID: 0},
	})
}

func CancelBooking(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.CancelBookingCourseResp{BookingID: 0},
	})
}

// Upload 上传课程
func Upload(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.UploadCourseResp{},
	})
}

// Bind 绑定用户课程
func Bind(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "success",
		"data": proto.BindCourseResp{},
	})
}
