package proto

// CommentCreateReq 发表评论请求
type CommentCreateReq struct {
	// ContentID 内容ID
	ContentID int `json:"content_id" form:"content_id"`
	// ContentType 内容类型
	ContentType uint8 `json:"content_type" form:"content_type"`
	// UserID 用户ID
	UserID int `json:"user_id" form:"user_id"`
	// CommentContent 评论内容
	CommentContent string `json:"comment_content" form:"comment_content"`
	// CommentPics 评论图片url, 多个用";"分割
	CommentPics string `json:"comment_pics" form:"comment_pics"`
	// CommentVideo 评论视频url
	CommentVideo string `json:"comment_video" form:"comment_video"`
	// CommentAudio 评论音频url
	CommentAudio string `json:"comment_audio" form:"comment_audio"`
	// PubTime 评论发表时间
	PubTime string `json:"pub_time" form:"pub_time"`
}

// CommentResp 评论回包
type CommentResp struct {
	CommentID int `json:"comment_id"`
}
