package proto

// ContentOperation 内容操作
type ContentOperation struct {
	// ContentID 内容ID
	ContentID int `form:"content_id" json:"content_id"`
	// UserID 用户ID
	UserID int `form:"user_id" json:"user_id"`
}

type ContentOperationResp struct {
}

// ContentViewReq 浏览内容请求
type ContentViewReq struct {
	// ContentOperation 嵌套ContentOperation 内容操作
	ContentOperation
}

// ContentCollectReq 收藏内容请求
type ContentCollectReq struct {
	// ContentOperation 嵌套ContentOperation 内容操作
	ContentOperation
}

// ContentCancelCollectReq 取消收藏内容请求
type ContentCancelCollectReq struct {
	// ContentOperation 嵌套ContentOperation 内容操作
	ContentOperation
}