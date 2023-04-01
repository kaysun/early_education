package proto

type BookingCourse struct {
	// UserID 用户ID
	UserID int `json:"user_id"`
	// CourseID 课程ID
	CourseID int `json:"course_id"`
}

// BookingCourseReq 约课请求
type BookingCourseReq struct {
	BookingCourse
}

// BookingCourseResp 约课回包
type BookingCourseResp struct {
	// BookingID 预约ID
	BookingID int `json:"booking_id"`
}

// CancelBookingCourseReq 取消约课请求
type CancelBookingCourseReq struct {
	BookingCourse
}

// CancelBookingCourseResp 取消约课回包
type CancelBookingCourseResp struct {
	// BookingID 预约ID
	BookingID int `json:"booking_id"`
}
