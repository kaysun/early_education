package proto

// BindCourseReq 绑定课程请求
type BindCourseReq struct {
	// Phone 手机号
	Phone int `json:"phone" form:"phone"`
	// CourseTimes 课时包总数
	CourseTimes int `json:"course_times" form:"course_times"`
}

// BindCourseResp 绑定课程回包
type BindCourseResp struct {
	// UserCourseID 用户课程id
	UserCourseID int `json:"user_course_id"`
}
