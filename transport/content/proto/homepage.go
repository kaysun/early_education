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
	Contents []ContentItem `json:"contents"`
}
