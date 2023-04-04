package proto

// Comment 评论
type Comment struct {
	// ContentID 内容ID
	ContentID int `json:"content_id" form:"content_id"`
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
}

// CommentResp 评论回包
type CommentResp struct {
	CommentID int `json:"comment_id"`
}

// ContentFollowAndPlayReq 跟玩内容请求
type ContentFollowAndPlayReq struct {
	// Comment 嵌套评论
	Comment
}
