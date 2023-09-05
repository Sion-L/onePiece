package model

import "time"

type User struct {
	ID         int
	Name       string
	Email      string
	DeptName   string
	CreateTime time.Time
	UpdateTime time.Time
}

func (u *User) TableName() string {
	return "user"
}
