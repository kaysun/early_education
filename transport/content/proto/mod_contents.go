package proto

// ModContentsReq 获取模块内容列表请求
type ModContentsReq struct {
	// ModID 模块ID
	ModID int `form:"mod_id"`
	// PageNo 当前页数
	PageNo int `form:"page_no"`
	// PageSize 每页数量
	PageSize int `form:"page_size"`
}

// ModContentsResp 获取模块内容列表回包
type ModContentsResp struct {
	// Total 模块内容总数
	Total int `json:"total"`
	// 模块信息
	Mod Mod `json:"mod"`
	// Contents 模块内容列表
	Contents []ContentItem `json:"contents"`
}

// ContentItem 内容项
type ContentItem struct {
	// ContentBase 嵌套内容
	ContentBase

	// ContentID 内容ID
	ContentID int `json:"content_id"`
	// CommentNum 评论数
	CommentNum int `json:"comment_num"`
	// CollectNum 收藏数
	CollectNum int `json:"collect_num"`
	// ViewsNum 浏览数
	ViewsNum int `json:"views_num"`
	// FollowAndPlayNum follow_and_play_num
	FollowAndPlayNum int `json:"follow_and_play_num"`
	// UpNum 点赞数
	UpNum int `json:"up_num"`
}
