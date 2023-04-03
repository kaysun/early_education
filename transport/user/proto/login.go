package proto

type LoginRequest struct {
	UserName string `json:"user_name" form:"user_name"`
	PassWord string `json:"pass_word" form:"pass_word"`
}

type LoginResp struct {
	// UserID 用户ID
	UserID int `json:"user_id"`
}
