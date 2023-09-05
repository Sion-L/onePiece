// package main
//
// import (
//
//	"fmt"
//	"github.com/Sion-L/onPiece/db"
//
// )
//
//	func main() {
//		conn, err := db.NewClientLdap()
//		if err != nil {
//			fmt.Print(err)
//		}
//		fmt.Print(conn)
//	}
package main

import (
	"fmt"
	pinyin "github.com/mozillazg/go-pinyin"
)

func main() {
	// 创建一个拼音转换器
	p := pinyin.Pinyin()

	// 将汉字转换为拼音，带声调
	hans := "你好，世界"
	result := p.Slug(hans, pinyin.WithTone)
	fmt.Println(result) // 输出: "nǐ-hǎo-shì-jiè"

	// 将汉字转换为拼音，不带声调
	result = p.Slug(hans, pinyin.WithoutTone)
	fmt.Println(result) // 输出: "ni-hao-shi-jie"
}
