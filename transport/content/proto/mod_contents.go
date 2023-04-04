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
	Contents []Content `json:"contents"`
}
