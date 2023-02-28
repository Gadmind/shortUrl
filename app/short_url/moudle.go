package short_url

import "sync"

var UrlInfoPool = &sync.Pool{New: func() interface{} { return new(UrlInfo) }}
var UrlDTOPool = &sync.Pool{New: func() interface{} { return new(UrlDTO) }}

// UrlInfo 链接地址信息
type UrlInfo struct {
	// 自增ID
	Id int `json:"id" db:"id"`
	// 短链接编码（62进制）
	RefCode string `json:"refCode" db:"ref_code"`
	// 原始URL
	OriginUrl string `json:"originUrl" db:"origin_url"`
	// 浏览次数
	PageView int `json:"pageView" db:"page_view"`
}

type UrlDTO struct {
	RefUrl string `json:"refUrl"`
}

func (u *UrlInfo) Free() {
	u.Id = 0
	u.RefCode = ""
	u.OriginUrl = ""
	u.PageView = 0
	UrlInfoPool.Put(u)
}

func (u *UrlDTO) Free() {
	u.RefUrl = ""
	UrlDTOPool.Put(u)
}
