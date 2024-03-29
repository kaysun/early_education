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
	// CommentID 评论ID
	CommentID int `json:"comment_id"`
}

// CommentReply 回复评论
type CommentReply struct {
	// RootID 根评论ID。若回复的是一级评论，则RootID=一级评论的CommentID；若回复的是二级评论，则RootID=二级评论的ParentID，即一级评论的CommentID。
	RootID int `json:"root_id" form:"root_id"`
	// ParentID 父评论ID。ParentID=要回复的CommentID，无论是一级评论，还是二级评论。注意：三级评论不开放回复入口。
	ParentID int `json:"parent_id" form:"parent_id"`
	// CommentCreateReq 嵌套创建评论请求
	CommentCreateReq
}

// CommentUpOrCancelUpReq 评论点赞/取消点赞请求
type CommentUpOrCancelUpReq struct {
	CommentID int `json:"comment_id" form:"comment_id"`
}

// CommentUpOrCancelUpResp 评论点赞/取消点赞回包
type CommentUpOrCancelUpResp struct {
}

// CommentDeleteReq 删除评论请求
type CommentDeleteReq struct {
	CommentID int `json:"comment_id" form:"comment_id"`
}

// CommentDeleteResp 删除评论请求回包
type CommentDeleteResp struct {
}
