package proto

type LoginRequest struct {
	UserName string `json:"user_name"`
	PassWord string `json:"pass_word"`
}
