package proto

// BookingCourseReq 约课请求
type BookingCourseReq struct {
	// UserID 用户ID
	UserID int `json:"user_id"`
	// CourseID 课程ID
	CourseID int `json:"course_id"`
}

// BookingCourseResp 约课回包
type BookingCourseResp struct {
	// BookingID 预约ID
	BookingID int `json:"booking_id"`
}
