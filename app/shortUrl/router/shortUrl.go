package router

import (
	"ShortUrl/app/shortUrl/models"
	"ShortUrl/common/binary"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"sync"
)

var startId = 3844
var lock sync.RWMutex
var urlMap = sync.Map{}
var referMap = sync.Map{}

// GenerateUrl 生成URL
func GenerateUrl(context *gin.Context) {
	url := models.UrlInfoPool.Get().(*models.UrlInfo)
	defer url.Free()
	body, _ := io.ReadAll(context.Request.Body)
	_ = json.Unmarshal(body, &url)
	value, ok := urlMap.Load(url.OriginUrl)
	if !ok {
		url.Id = generateId()
		url.RefCode = binary.ConversionWithBinary(int64(url.Id), 62)
		urlMap.Store(url.OriginUrl, url.RefCode)
		referMap.Store(url.RefCode, url.OriginUrl)
		value = url.RefCode
	}
	context.JSON(http.StatusOK, gin.H{
		"code":    http.StatusOK,
		"success": true,
		"message": "生成成功",
		"data":    value,
	})
}

func Redirect(context *gin.Context) {
	path, _ := context.Params.Get("path")
	value, ok := referMap.Load(path)
	if !ok {
		context.JSON(200, gin.H{
			"status":  404,
			"message": "404 Not Found" + context.Request.Method + " " + context.Request.RequestURI,
		})
	} else {
		n := fmt.Sprintf("%s", value)
		context.Redirect(http.StatusFound, n)
	}
}

func generateId() int {
	lock.Lock()
	defer lock.Unlock()
	startId += 1
	return startId
}
