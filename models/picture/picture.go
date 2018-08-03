package models

import (
	"time"
)

type Picture struct {
	// 图片名
	Name string
	// 是否公开
	IsPublic bool
	// 上传时间
	Time time.Time
}

