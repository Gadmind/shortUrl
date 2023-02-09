package router

import (
	"ShortUrl/app/shortUrl/models"
	"ShortUrl/common/binary"
	"ShortUrl/common/constant"
	"ShortUrl/common/rand"
	"ShortUrl/common/reids"
	"ShortUrl/common/resp"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
)

// GenerateUrl 生成短连接URL
func GenerateUrl(context *gin.Context) {
	urlInfo := models.UrlInfoPool.Get().(*models.UrlInfo)
	defer urlInfo.Free()
	body, _ := io.ReadAll(context.Request.Body)
	_ = json.Unmarshal(body, &urlInfo)

	originUrl := urlInfo.OriginUrl
	refCode := reids.Get(constant.SOSPrefix + originUrl)
	if refCode == "" {
		randomId := rand.GetRandomId()
		refCode = binary.ConversionWithBinary(int64(randomId), 62)
		reids.Set(refCode, constant.SOSPrefix+originUrl, 0)
		reids.Set(originUrl, constant.SSOPrefix+refCode, 0)
		// TODO 入库保存
	}
	context.JSON(resp.Success("生成成功", refCode))
}

// Redirect 短链接重定向
func Redirect(context *gin.Context) {
	path, _ := context.Params.Get("path")
	value := reids.Get(constant.SSOPrefix + path)
	if value == "" {
		context.JSON(resp.ReNoData(http.StatusNotFound, false, "404 Not Found"+context.Request.Method+" "+context.Request.RequestURI))
	}
	n := fmt.Sprintf("%s", value)
	context.Redirect(http.StatusFound, n)
}
