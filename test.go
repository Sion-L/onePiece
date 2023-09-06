package main

import (
	"fmt"
	"github.com/Sion-L/onePiece/dao"
)

func main() {
	err := dao.LdapDeleteUser("萧何")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("删除成功")

}

// TODO:
// 1. 添加用户后往数据库同步  完成
// 2. 删除用户后从数据库删除
// 3. ldap支持更新密码
// 4. 用户更新组织后数据库也更新
