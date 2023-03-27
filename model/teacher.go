package model

// Teacher 教师
type Teacher struct {
	// Name 名字
	Name string `json:"name"`
	// Age 年龄
	Age uint8 `json:"age"`
	// Description 描述
	Description string `json:"description"`
	// CourseIDs 教授课程ID列表
	CourseIDs []uint32 `json:"-"`
}
