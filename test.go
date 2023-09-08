package main

import (
	"fmt"
	"github.com/Sion-L/onePiece/dao"
	"github.com/Sion-L/onePiece/db"
)

func init() {
	db.InitLdap()
	db.InitMySQLDB()
}

func main() {
	res := dao.LoginForLdap("lilang", "123456")
	fmt.Print(res)
}
