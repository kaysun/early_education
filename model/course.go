package model

import "time"

// Course 课程
type Course struct {
	// CourseID 课程id
	CourseID uint32 `gorm:"column:course_id"`
	// ClassName 课程名字
	ClassName string `gorm:"column:name"`
	// ClassTime 上课时间
	ClassTime time.Time `gorm:"column:class_time"`
	// TeacherID 教师ID
	TeacherID int `gorm:"column:teacher_id"`
	// MaxStudents 课程最多容纳的学生数量
	MaxStudents int `gorm:"column:max_students"`
	// ScheduledStudents 已经约课的学生数
	ScheduledStudents int `gorm:"column:scheduled_students"`
	// CourseStatus 课程状态：1-有空闲 2-约满
	CourseStatus int `gorm:"column:course_status"`
	// OpenState 开放状态 0-未开放 1-开放中
	OpenState uint8 `gorm:"column:open_state"`
	// CreateTime 创建时间
	CreateTime time.Time `gorm:"column:create_time"`
	// UpdateTime 修改时间
	UpdateTime time.Time `gorm:"column:update_time"`
}
