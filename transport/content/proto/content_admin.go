package proto

// ContentUploadReq 内容上传请求
type ContentUploadReq struct {
	// ContentBase 嵌套内容
	ContentBase
}

// ContentUploadResp 内容上传回包
type ContentUploadResp struct {
	ContentID int `json:"content_id"`
}
