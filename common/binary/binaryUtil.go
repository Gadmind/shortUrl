package binary

import (
	"math"
	"strings"
)

var char = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// ConversionWithBinary 10进制转换为其他进制
func ConversionWithBinary(num int64, bin int64) string {
	var bytes []byte
	for num > 0 {
		bytes = append(bytes, char[num%bin])
		num /= bin
	}
	reverse(bytes)
	return string(bytes)
}

// OtherBinaryConversion 其他进制转换为10进制
func OtherBinaryConversion(str string, bin float64) int64 {
	var num int64
	l := len(str)
	for i := 0; i < l; i++ {
		pos := strings.IndexByte(char, str[i])
		num += int64(math.Pow(bin, float64(l-i-1)*float64(pos)))
	}
	return num
}

// 数组反转
func reverse(bytes []byte) {
	for left, right := 0, len(bytes)-1; left < right; left, right = left+1, right-1 {
		bytes[left], bytes[right] = bytes[right], bytes[left]
	}
}
