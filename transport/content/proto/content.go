package proto

// ContentBase 内容
type ContentBase struct {
	// ChannelID 频道ID
	ChannelID int `json:"channel_id" form:"channel_id"`
	// ModID 模块ID
	ModID int `json:"mod_id" form:"mod_id"`
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
	// CreateTime 创建时间
	CreateTime string `json:"create_time"`
}
