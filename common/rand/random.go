package rand

import (
	"ShortUrl/common/constant"
	"ShortUrl/common/reids"
	rand2 "crypto/rand"
	"math/big"
	"sync"
)

var (
	lock    sync.RWMutex
	startId = 3844
)

// GetRandomId 获取随机ID
func GetRandomId() int {
	rnd, _ := rand2.Int(rand2.Reader, big.NewInt(constant.MaxRandomRange))
	for true {
		if rnd.Int64() >= constant.MinRandomRange {
			isExist := reids.Get(constant.RandPrefix + rnd.String())
			if isExist == "" {
				reids.Set(rnd.String(), constant.RandPrefix+rnd.String(), 0)
			}
			return int(rnd.Int64())
		} else {
			rnd, _ = rand2.Int(rand2.Reader, big.NewInt(constant.MaxRandomRange))
			continue
		}
	}
	return 0
}

// GenerateId 获取ID
func GenerateId() int {
	lock.Lock()
	defer lock.Unlock()
	startId += 1
	return startId
}
