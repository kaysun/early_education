package proto

// TopicsReq 专题列表请求
type TopicsReq struct {
	UserID int `json:"user_id"`
}

// TopicsResp 专题列表回包
type TopicsResp struct {
	// OfficialCreationTopics 官方创建专题列表
	OfficialCreationTopics []Topic `json:"official_creation_topics"`
	// OfficialRecommendationTopics 官方推荐专题列表
	OfficialRecommendationTopics []Topic `json:"official_recommendation_topics"`
	// UserCreationTopics 用户自己创建的专题列表
	UserCreationTopics []Topic `json:"user_creation_topics"`
	// UserParticipatingTopics 用户参与的专题列表
	UserParticipatingTopics []Topic `json:"user_participating_topics"`
}

// Topic 专题
type Topic struct {
	// TopicID 专题iD
	TopicID int `json:"topic_id"`
	// TopicName 专题名字
	TopicName string `json:"topic_name"`
	// TopicDesc 专题描述
	TopicDesc string `json:"topic_desc"`
	// TopicCover 专题封面图
	TopicCover string `json:"topic_cover"`
	// TopicType 1-官方创建 2-官方推荐 3-自己创建 4-他人创建
	TopicType uint8 `json:"topic_type"`
	// CreatorUserID 创建人 官方：0； 个人创建：user_id
	CreatorUserID int `json:"creator_uid"`
	// CommentNum 评论数
	CommentNum int `json:"comment_num"`
	// CreateTime 创建时间
	CreateTime string `json:"create_time"`
}

// TopicCreateReq 创建专题请求
type TopicCreateReq struct {
	// TopicName 专题名字
	TopicName string `json:"topic_name"`
	// TopicDesc 专题描述
	TopicDesc string `json:"topic_desc"`
	// TopicCover 专题封面图
	TopicCover string `json:"topic_cover"`
	// UserID 创建人用户ID
	UserID int `json:"user_id"`
}

// TopicCreateResp 创建专题回包
type TopicCreateResp struct {
	// TopicID 专题ID
	TopicID int `json:"topic_id"`
}
