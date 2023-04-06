package proto

// CommentListReq 分页获取内容下的评论列表请求
type CommentListReq struct {
	// ContentID 内容ID
	ContentID int `form:"content_id"`
	// PageNo 当前页数
	PageNo int `form:"page_no"`
	// PageSize 每页数量
	PageSize int `form:"page_size"`
}

// CommentListResp 分页获取内容下的评论列表回包
type CommentListResp struct {
	// Total 内容评论总数
	Total int `json:"total"`
	// Contents 模块内容列表
	Comments []CommentItem `json:"comments"`
}

// CommentItem 评论项
type CommentItem struct {
	// ContentID 内容ID
	ContentID int `json:"content_id"`

	// UserID 用户ID
	UserID int `json:"user_id"`
	// UserNick 用户昵称
	UserNick string `json:"user_nick"`
	// UserHead 用户头像
	UserHead string `json:"user_head"`

	// CommentContent 评论内容
	CommentContent string `json:"comment_content"`
	// CommentPics 评论图片url, 多个用";"分割
	CommentPics string `json:"comment_pics"`
	// CommentVideo 评论视频url
	CommentVideo string `json:"comment_video"`
	// CommentAudio 评论音频url
	CommentAudio string `json:"comment_audio"`
	// PubTime 评论发表时间
	PubTime string `json:"pub_time"`
}
