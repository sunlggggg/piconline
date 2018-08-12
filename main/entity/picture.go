package entity

type Picture struct {
	// ID
	Id uint64 `json:"id"`
	// 图片名
	Name string `json:"name"`
	// 是否公开
	IsPublic bool `json:"is_public"`
	// 上传时间
	Time uint64 `json:"time"`
	// 描述信息
	Desc string `json:"description"`
	// 拥有者
	UserId uint64 `json:"user_id"`
}
