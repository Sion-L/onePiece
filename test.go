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
	ok := dao.LoginForLdap("lilang", "123456")
	fmt.Println(ok)
}
