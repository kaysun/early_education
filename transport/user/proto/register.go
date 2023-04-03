package proto

// RegisterReq 注册请求
type RegisterReq struct {
	// Name 用户名，全局唯一
	Name string `json:"name" form:"name"`
	// Password 密码，存储密文（这里前端就应该传入密文，但后台依然可以按照做一层md5加密，加一层保护，避免前端传明文，后端要假设前端不可信）
	Password string `json:"password" form:"password"`
	// Nick 昵称
	Nick string `json:"nick" form:"nick"`
	// Birthday 出生日期，用于判定
	Birthday string `json:"birthday" form:"birthday"`
}

type RegisterResp struct {
	// UserID 用户ID
	UserID int `json:"user_id"`
}
