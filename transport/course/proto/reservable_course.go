package proto

// ReservableListReq 获取可约课程列表请求
type ReservableListReq struct {
	// UserID 用户ID
	UserID int `json:"user_id"`
}

// ReservableListResp 获取可约课程列表回包
type ReservableListResp struct {
	ReservableCourses []ReservableCourse `json:"reservable_courses"`
}
