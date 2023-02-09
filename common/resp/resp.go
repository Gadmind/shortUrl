package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 响应成功JSON数据
func Success(msg string, data any) (int, gin.H) {
	return ReCustom(http.StatusOK, true, msg, data)
}

// Fail 响应失败JSON数据
func Fail(msg string, data any) (int, gin.H) {
	return ReCustom(http.StatusInternalServerError, false, msg, data)
}

// ReNoData 无返回数据响应
func ReNoData(code int, success bool, msg string) (int, gin.H) {
	return ReCustom(code, success, msg, nil)
}

// ReCustom 自定义JSON响应
func ReCustom(code int, success bool, msg string, data any) (int, gin.H) {
	return code, gin.H{
		"code":    code,
		"success": success,
		"message": msg,
		"data":    data,
	}
}
