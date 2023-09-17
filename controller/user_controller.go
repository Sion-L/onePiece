package controller

import (
	"fmt"
	"github.com/Sion-L/onePiece/dao"
	"github.com/Sion-L/onePiece/middleware"
	"github.com/Sion-L/onePiece/model"
	"github.com/Sion-L/onePiece/types"
	"github.com/Sion-L/onePiece/utils"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const DefaultDepartName = "employee"

func AddUserForLdap(c *gin.Context) {
	var form types.LdapUser
	if err := c.Bind(&form); err != nil {
		types.Fail(c, http.StatusOK, "绑定数据错误")
		return
	}
	name, err := dao.FilterUser(form.Cn)
	if err != nil {
		types.Fail(c, http.StatusOK, err.Error())
		return
	}
	err = dao.AddUser(name[0], form.Cn, form.Password)
	if err != nil {
		types.FailF(c, http.StatusOK, fmt.Sprintf("添加用户%s失败", name[0]), err.Error())
		return
	}
	err = AddUserForDB(form.Cn, name[0])
	if err != nil {
		types.FailF(c, http.StatusOK, fmt.Sprintf("用户%s信息入库失败", name[0]), err.Error())
		return
	}
	types.Success(c, fmt.Sprintf("添加用户%s成功", name[0]))
}

func AddUserForDB(cn, sn string) error {
	user := &model.User{
		ID:         utils.GenerateNumericUserID(11),
		Cn:         cn,
		En:         sn,
		Email:      fmt.Sprintf("%s@lang.com", sn),
		DeptName:   DefaultDepartName,
		Role:       "",
		Business:   "",
		CreateTime: time.Now().Format("2006-01-02 15:04:05"),
		UpdateTime: time.Now().Format("2006-01-02 15:04:05"),
	}
	err := dao.InsertUserMany(user)
	if err != nil {
		return err
	}
	return nil
}

// 删除用户
func DeleteUser(c *gin.Context) {

	var form types.LdapDeleteUser
	if err := c.Bind(&form); err != nil {
		types.Fail(c, http.StatusOK, "绑定数据错误")
		return
	}
	err := dao.LdapDeleteUser(form.Id, form.Cn)
	if err != nil {
		types.FailF(c, http.StatusOK, fmt.Sprintf("删除用户%s失败", form.Cn), err.Error())
		return
	}

	types.Success(c, fmt.Sprintf("删除用户%s成功", form.Cn))
}

// 修改密码
func ResetPassword(c *gin.Context) {
	var form types.LdapResetPass
	if err := c.Bind(&form); err != nil {
		types.Fail(c, http.StatusOK, "绑定数据错误")
		return
	}
	sn, _ := dao.FindUserByLdap(form.Cn)

	err := dao.LdapResetPassword(sn, form.Password)
	if err != nil {
		types.FailF(c, http.StatusOK, fmt.Sprintf("修改用户%s密码失败", sn), err.Error())
		return
	}
	types.Success(c, fmt.Sprintf("修改用户%s密码成功", sn))
}

// 登陆
func Login(c *gin.Context) {
	var form types.LoginUser
	if err := c.Bind(&form); err != nil {
		types.Fail(c, http.StatusOK, "绑定数据错误")
		return
	}

	if ok := dao.LoginForLdap(form.En, form.Password); !ok {
		types.Fail(c, http.StatusUnauthorized, "用户名或密码不正确")
		return
	}

	tokenString, err := middleware.ReleaseToken(form.En, form.Password)
	if err != nil {
		types.Fail(c, http.StatusInternalServerError, "Error generating token")
		return
	}
	c.Writer.Header().Set("One-Piece", tokenString)
	color.Yellow(tokenString)

	types.Success(c, "登陆成功")
}

func FindAllUser(c *gin.Context) {
	users, err := dao.FindAllUser()
	if err != nil {
		types.Fail(c, http.StatusOK, err.Error())
		return
	}
	types.Success(c, users)
}