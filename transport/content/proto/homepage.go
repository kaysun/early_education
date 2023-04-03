package proto

// HomepageReq 首页内容请求
type HomepageReq struct {
	UserID int `form:"user_id"`
}

// HomepageResp 首页内容回包
type HomepageResp struct {
	// Mods 模块列表
	Mods []Mod `json:"mods"`
}

// Mod 模块
type Mod struct {
	// ModID 模块ID
	ModID int `json:"mod_id"`
	// ModName 模块名字
	ModName string `json:"mod_name"`
	// ModType 模块类型
	ModType uint8 `json:"mod_type"`
	// Contents 内容列表
	Contents []Content `json:"contents"`
}

// Content 内容
type Content struct {
	// ContentID 内容ID
	ContentID int `json:"content_id"`
	// ContentType 内容类型
	ContentType uint8 `json:"content_type"`
	// PrimaryClassification 一级分类
	PrimaryClassification string `json:"primary_classification"`
	// SecondaryClassification 二级分类
	SecondaryClassification string `json:"secondary_classification"`
	// CoverImgURL 封面图
	CoverImgURL string `json:"cover_img"`
	// Title 标题
	Title string `json:"title"`
	// Description 描述
	Description string `json:"description"`
	// ContentURL 内容存储位置
	ContentURL string `json:"content_url"`
	// Version 版本号
	Version int `json:"version"`
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
	// CreateTime 创建时间
	CreateTime string `json:"create_time"`
}
