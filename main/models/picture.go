package models

import (
	"time"
)

type Picture struct {
	// 图片名
	Name string `xml:"name"`
	// 是否公开
	IsPublic bool `xml:"is_public"`
	// 上传时间
	Time time.Time `xml:"time"`
}

