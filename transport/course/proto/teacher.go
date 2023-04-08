package proto

type TeacherListReq struct {
}

// TeacherListResp 教师列表回包
type TeacherListResp struct {
	Teachers []Teacher `json:"teachers"`
}

// Teacher 教师
type Teacher struct {
	// TeacherID 教师ID，指定了教师ID，则以教师ID为准；未指定教师ID，则根据教师信息创建教师信息，并指定给课程
	TeacherID int `json:"teacher_id" form:"teacher_id"`
	// TeacherName 教师名字
	TeacherName string `json:"teacher_name" form:"teacher_name"`
	// TeacherAge 教师年龄
	TeacherAge uint8 `json:"teacher_age" form:"teacher_age"`
	// TeacherDesc 教师描述
	TeacherDesc string `json:"teacher_desc" form:"teacher_desc"`
}
