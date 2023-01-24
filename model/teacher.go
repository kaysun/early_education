package model

// Teacher 教师
type Teacher struct {
	// Name 名字
	Name string `json:"name"`
	// Age 年龄
	Age uint8 `json:"age"`
	// Area 地域
	Area string `json:"area"`
	// CourseIDs 教授课程列表
	CourseIDs []uint32 `json:"-"`
}
