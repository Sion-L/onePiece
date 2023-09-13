package model

type User struct {
	ID         string `json:"id" binding:"required"`
	Cn         string `json:"cn" binding:"required"`
	En         string `json:"en" binding:"required"`
	Email      string `json:"email"`
	DeptName   string `json:"dept_name"`
	Role       string `json:"role"`
	Business   string `json:"business"`
	CreateTime string `json:"create_time"`
	UpdateTime string `json:"update_time"`
}

func (u *User) TableName() string {
	return "user"
}
