package dto

import "sync"

var UrlDTOPool = &sync.Pool{New: func() interface{} {
	return new(UrlDTO)
}}

type UrlDTO struct {
	RefUrl string `json:"refUrl"`
}

func (u *UrlDTO) Free() {
	u.RefUrl = ""
	UrlDTOPool.Put(u)
}
