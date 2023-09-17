package main

import (
	"fmt"
	"github.com/Sion-L/onePiece/db"
	"github.com/Sion-L/onePiece/middleware"
)

func init() {
	db.InitLdap()
	db.InitMySQLDB()
}

func main() {
	user, err := middleware.ParseToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImxpbGFuZyIsInBhc3N3b3JkIjoiNzI3MzcxMjczMSIsImV4cCI6MTY5NTA1MTEwOCwiaWF0IjoxNjk0OTY0NzA4LCJpc3MiOiJsYW5nIiwic3ViIjoidXNlciB0b2tlbiJ9.Ui2x1Uz62IJTPGu2mRhClh6lRC0KWaBvubOPA6WMmOc")
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(user)
}
