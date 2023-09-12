package controller

import (
	"fmt"
	"github.com/Sion-L/onePiece/dao"
	"github.com/Sion-L/onePiece/model"
	"github.com/Sion-L/onePiece/types"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func AddUserForLdap(c *gin.Context) {
	var form types.LdapUser
	if err := c.Bind(&form); err != nil {
		types.Fail(c, http.StatusOK, "绑定数据错误")
		return
	}
	name, err := dao.FilterUser(form.CN)
	if err != nil {
		types.Fail(c, http.StatusOK, err.Error())
		return
	}
	err = dao.AddUser(form.OU, name[0], form.CN, form.Pass)
	if err != nil {
		types.Fail(c, http.StatusOK, fmt.Sprintf("添加用户 %s失败:%s", name[0], err.Error()))
		return
	}
	err = AddUserForDB(form.OU, form.CN, name[0])
	if err != nil {
		types.Fail(c, http.StatusOK, fmt.Sprintf("用户%s信息入库失败: %s", name[0], err.Error()))
		return
	}
	types.Success(c, fmt.Sprintf("添加用户%s成功", name[0]))
}

func AddUserForDB(ou, cn, sn string) error {
	user := &model.User{
		ID:         types.GenerateUnique([]string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}),
		UserName:   cn,
		Email:      fmt.Sprintf("%s@lang.com", sn),
		DeptName:   ou,
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

	err := dao.LdapDeleteUser(form.CN)
	if err != nil {
		types.Fail(c, http.StatusOK, fmt.Sprintf("删除用户%s失败, %s", form.CN, err.Error()))
		return
	}
	types.Success(c, fmt.Sprintf("删除用户%s成功", form.CN))
}

// 修改密码
func ResetPassword(c *gin.Context) {
	var form types.LdapResetPass
	if err := c.Bind(&form); err != nil {
		types.Fail(c, http.StatusOK, "绑定数据错误")
		return
	}
	sn, _ := dao.FindUserByLdap(form.CN)

	err := dao.LdapResetPassword(sn, form.Password)
	if err != nil {
		types.Fail(c, http.StatusOK, fmt.Sprintf("修改用户%s密码失败: %s", sn, err))
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
	if ok := dao.LoginForLdap(form.UserName, form.Password); !ok {
		types.Fail(c, http.StatusUnauthorized, "用户名或密码不正确")
		return
	}
	types.Success(c, "登陆成功")
}
