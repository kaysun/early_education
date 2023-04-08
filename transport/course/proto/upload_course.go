package proto

// UploadCourseReq 上传课程请求
type UploadCourseReq struct {
	// CourseName 课程名字
	CourseName string `json:"course_name" form:"course_name"`
	// ClassTime
	ClassTime string `json:"class_time" form:"class_time"`
	// MaxStudents 课程最多容纳的学生数量
	MaxStudents int `json:"max_students" form:"max_students"`
	// ScheduledStudents 已经约课的学生数
	ScheduledStudents int `json:"scheduled_students" form:"scheduled_students"`

	// 嵌套教师，指定了教师ID，则以教师ID为准；未指定教师ID，则根据教师信息创建教师信息，并指定给课程
	Teacher
	// AgeStages 课程适用年龄列表
	AgeStages []uint8 `json:"age_stage"`
}

// UploadCourseResp 上传课程回包
type UploadCourseResp struct {
	CourseID int `json:"course_id"`
}
