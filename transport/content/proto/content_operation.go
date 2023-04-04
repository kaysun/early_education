package proto

// ContentOperation 内容操作
type ContentOperation struct {
	// ContentID 内容ID
	ContentID int `json:"content_id" form:"content_id"`
	// UserID 用户ID
	UserID int `json:"user_id" form:"user_id"`
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

// ContentUpReq 点赞内容
type ContentUpReq struct {
	// ContentOperation 嵌套ContentOperation 内容操作
	ContentOperation
}

// ContentCancelUpReq 取消点赞内容
type ContentCancelUpReq struct {
	// ContentOperation 嵌套ContentOperation 内容操作
	ContentOperation
}
