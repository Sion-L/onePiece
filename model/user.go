package model

type User struct {
	ID         int    `json:"id" binding:"required"`
	Name       string `json:"name" binding:"required"`
	Email      string `json:"email" binding:"required"`
	DeptName   string `json:"dept_name" binding:"required"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}
