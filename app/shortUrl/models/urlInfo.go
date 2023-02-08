package models

import (
	"sync"
)

var UrlInfoPool = &sync.Pool{New: func() interface{} { return new(UrlInfo) }}

// UrlInfo 链接地址信息
type UrlInfo struct {
	// 自增ID
	Id int `json:"id"`
	// 短链接编码（62进制）
	RefCode string `json:"refCode"`
	// 原始URL
	OriginUrl string `json:"originUrl"`
	// 浏览次数
	PageView int `json:"pageView"`
	// TODO 其他字段待添加
}

func (u *UrlInfo) Free() {
	u.Id = 0
	u.RefCode = ""
	u.OriginUrl = ""
	u.PageView = 0
	UrlInfoPool.Put(u)
}
