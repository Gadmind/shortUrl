package router

import (
	"ShortUrl/app/shortUrl/models"
	"ShortUrl/app/shortUrl/models/dto"
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
	"strings"
)

// GenerateUrl 生成短连接URL
func GenerateUrl(context *gin.Context) {
	urlInfo := models.UrlInfoPool.Get().(*models.UrlInfo)
	defer urlInfo.Free()
	body, _ := io.ReadAll(context.Request.Body)
	_ = json.Unmarshal(body, &urlInfo)

	originUrl := urlInfo.OriginUrl
	if !(strings.HasPrefix(originUrl, "https://") || strings.HasPrefix(originUrl, "http://")) {
		context.JSON(resp.ReNoData(200, false, "链接格式不正确"))
		return
	}

	refCode := reids.Get(constant.SOSPrefix + originUrl)
	if refCode == "" {
		randomId := rand.GetRandomId()
		refCode = binary.ConversionWithBinary(int64(randomId), 62)
		reids.Set(refCode, constant.SOSPrefix+originUrl, 0)
		reids.Set(originUrl, constant.SSOPrefix+refCode, 0)
		// TODO 入库保存
	}
	context.JSON(resp.Success("生成成功", context.Request.Host+"/"+refCode))
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

// AnalysisUrl 解析短URL
func AnalysisUrl(context *gin.Context) {
	url := dto.UrlDTOPool.Get().(*dto.UrlDTO)
	defer url.Free()
	body, _ := io.ReadAll(context.Request.Body)
	_ = json.Unmarshal(body, &url)
	refUrl := url.RefUrl
	if !(strings.HasPrefix(refUrl, "https://") || strings.HasPrefix(refUrl, "http://")) {
		context.JSON(resp.ReNoData(200, false, "链接格式不正确"))
		return
	}
	urls := strings.Split(refUrl, "/")
	refCode := urls[len(urls)-1]
	originUrl := reids.Get(constant.SSOPrefix + refCode)
	if originUrl == "" {
		context.JSON(resp.SuccessNoData("该链接不存在或已失效"))
	}
	context.JSON(resp.Success("解析成功", originUrl))
}

// DeleteUrl 删除URL
func DeleteUrl(context *gin.Context) {
	url := dto.UrlDTOPool.Get().(*dto.UrlDTO)
	defer url.Free()
	body, _ := io.ReadAll(context.Request.Body)
	_ = json.Unmarshal(body, &url)
	refUrl := url.RefUrl
	if !(strings.HasPrefix(refUrl, "https://") || strings.HasPrefix(refUrl, "http://")) {
		context.JSON(resp.ReNoData(200, false, "链接格式不正确"))
		return
	}
	urls := strings.Split(refUrl, "/")
	refCode := urls[len(urls)-1]
	originUrl := reids.Get(constant.SSOPrefix + refCode)
	if originUrl == "" {
		context.JSON(resp.SuccessNoData("该链接不存在或已删除"))
	}
	reids.Del(constant.SSOPrefix + refCode)
	reids.Del(constant.SOSPrefix + originUrl)
	context.JSON(resp.SuccessNoData("该链接已删除"))
}
