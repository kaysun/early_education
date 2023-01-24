package model

// Course 课程
type Course struct {
	// CourseID 课程id
	CourseID uint32 `json:"course_id"`
	// Mechanism 机构
	Mechanism string `json:"mechanism"`
	// ClassTime 课程时间
	ClassTime string `json:"class_time"`
	// ClassName 课程名字
	ClassName string `json:"class_name"`
	// Teacher 教师
	Teacher `json:"teacher"`
	// OpenState 开放状态 0-未开放 1-开放中
	OpenState uint8 `json:"open_state"`
}
