package dao

import (
	"fmt"
	"github.com/Sion-L/onePiece/db"
	"github.com/Sion-L/onePiece/model"
)

func InsertUserMany(user *model.User) error {
	result := db.Conn.Create(user)
	if result.Error != nil {
		return fmt.Errorf("error inserting records: %v", result.Error)
	}

	return nil
}

// 查找用户是否存在
func FindUserByName(en string) (*model.User, error) {
	var user model.User
	result := db.Conn.Where("en = ?", en).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

// 删除用户
func DeleteUserByUserId(id string) error {

	result := db.Conn.Where("id = ?", id).Delete(&model.User{})
	if result.Error != nil {
		return result.Error
	}

	return nil
}

// 查询所有用户
func FindAllUser() ([]*model.User, error) {
	var users []*model.User
	result := db.Conn.Find(&users)
	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}