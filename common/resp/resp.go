package resp

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// Success 响应成功JSON数据
func Success(msg string, data any) (int, gin.H) {
	return ReCustom(http.StatusOK, true, msg, data)
}

// SuccessNoData 响应成功消息
func SuccessNoData(msg string) (int, gin.H) {
	return Success(msg, nil)
}

// Fail 响应失败JSON数据
func Fail(msg string, data any) (int, gin.H) {
	return ReCustom(http.StatusInternalServerError, false, msg, data)
}

// FailNoData 响应失败消息
func FailNoData(msg string) (int, gin.H) {
	return Fail(msg, nil)
}

// ReNoData 无返回数据响应
func ReNoData(code int, success bool, msg string) (int, gin.H) {
	return ReCustom(code, success, msg, nil)
}

// ReCustom 自定义JSON响应
func ReCustom(code int, success bool, msg string, data any) (int, gin.H) {
	result := gin.H{
		"code":    code,
		"success": success,
		"message": msg,
	}
	if data != nil {
		result["data"] = data
	}
	return code, result
}
