package utils

import (
	"math/rand"
	"time"
)

// 生成userid
func GenerateNumericUserID(length int) string {
	// 初始化随机数生成器
	rand.Seed(time.Now().UnixNano())

	// 生成随机数字字符串
	digits := "0123456789"
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = digits[rand.Intn(len(digits))]
	}

	return string(result)
}
