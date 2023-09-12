package main

import (
	"fmt"
	"github.com/Sion-L/onePiece/types"
)

//func init() {
//	db.InitLdap()
//	db.InitMySQLDB()
//}

func main() {
	//res := dao.LoginForLdap("lilang", "123456")
	//fmt.Print(res)

	pass := types.GenerateUnique([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"})
	fmt.Println(pass)
}
