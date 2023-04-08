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

	// TeacherID 教师ID，指定了教师ID，则以教师ID为准；未指定教师ID，则根据教师信息创建教师信息，并指定给课程
	TeacherID int `json:"teacher_id" form:"teacher_id"`
	// TeacherName 教师名字
	TeacherName string `json:"teacher_name" form:"teacher_name"`
	// TeacherAge 教师年龄
	TeacherAge uint8 `json:"teacher_age" form:"teacher_age"`
	// TeacherDesc 教师描述
	TeacherDesc string `json:"teacher_desc" form:"teacher_desc"`
}

// UploadCourseResp 上传课程回包
type UploadCourseResp struct {
	CourseID int `json:"course_id"`
}
