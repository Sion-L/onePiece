package dao

import (
	"fmt"
	"github.com/Sion-L/onePiece/db"
	"github.com/Sion-L/onePiece/model"
)

func InsertUserMany(user *model.User) error {
	result := db.Conn.Create(user)
	if result.Error != nil {
		return fmt.Errorf("error inserting records: %s", result.Error.Error())
	}

	return nil
}

// 查找用户是否存在
func FindUserByName(name string) (*model.User, error) {
	var user model.User
	result := db.Conn.Where("name = ?", name).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// 删除用户
func DeleteUserByName(name string) error {

	result := db.Conn.Where("user_name = ?", name).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}
