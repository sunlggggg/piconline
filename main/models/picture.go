package models

import (
	"time"
)

type Picture struct {
	// 图片名
	Name string `json:"name"`
	// 是否公开
	IsPublic bool `json:"is_public"`
	// 上传时间
	Time time.Time `json:"time"`
}

