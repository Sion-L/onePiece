package dao

import (
	"fmt"
	"github.com/Sion-L/onePiece/db"
	"github.com/Sion-L/onePiece/model"
)

func InsertUserMany(user *model.User) error {
	db, err := db.NewClientDB()
	if err != nil {
		return err
	}

	result := db.Create(user)
	if result.Error != nil {
		return fmt.Errorf("error inserting records: %s", result.Error.Error())
	}

	return nil
}
