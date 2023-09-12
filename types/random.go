package types

import (
	"fmt"
	"math/rand"
	"time"
)

func GenerateUnique(torrent []string) string {
	// 设置随机数生成器的种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个包含0到9的数字的切片
	//digits := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}

	// 打乱切片中的数字顺序
	rand.Shuffle(len(torrent), func(i, j int) {
		torrent[i], torrent[j] = torrent[j], torrent[i]
	})

	// 从打乱后的切片中选择前八个数字
	result := torrent[:8]

	// 将数字拼接成字符串
	userID := ""
	for _, digit := range result {
		userID += fmt.Sprintf("%s", digit)
	}

	return userID
}
