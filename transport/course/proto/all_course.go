package proto

import "time"

// GetAllCoursesReq 获取全量课程请求（默认返回未来7天的课程）
type GetAllCoursesReq struct {
	// UserID 用户ID
	UserID int `form:"user_id"`
}

// GetAllCoursesResp 获取全量课程回包
type GetAllCoursesResp struct {
	Courses []Course `json:"courses"`
}

// Course 课程
type Course struct {
	// CourseID 课程id
	CourseID int `json:"course_id"`
	// CourseName 课程名字
	CourseName string `json:"course_name"`
	// ClassTime 上课时间
	ClassTime time.Time `json:"class_time"`
	// MaxStudents 课程最多容纳的学生数量
	MaxStudents int `json:"max_students"`
	// ScheduledStudents 已经约课的学生数
	ScheduledStudents int `json:"scheduled_students"`
	// CourseStatus 课程状态：1-有空闲 2-约满
	CourseStatus int `json:"course_status"`

	// TeacherID 教师ID
	TeacherID int `json:"teacher_id"`
	// TeacherName 教师名字
	TeacherName int `json:"teacher_name"`
	// TeacherAge 教师年龄
	TeacherAge uint8 `json:"teacher_age"`
	// TeacherDesc 教师优势描述
	TeacherDesc string `json:"teacher_desc"`
}

// ReservableCourse 可约的课程
type ReservableCourse struct {
	Course
	// IsScheduled 是否已约
	IsScheduled bool `json:"is_scheduled"`
}
