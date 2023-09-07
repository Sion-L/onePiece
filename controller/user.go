// package controller
//
// import (
//
//	"github.com/Sion-L/gin-demo/models"
//	"github.com/gin-gonic/gin"
//
// )
//
//	func GetProject(c *gin.Context) {
//		var list []string
//		client := models.GitlabOptions{
//			Url:      "http://gitlab.firecloud.wan/",
//			Username: "lilang",
//			Password: "ll772576",
//		}
//
//		groups, err := client.GetGroups()
//		if err != nil {
//			return
//		}
//		for _, v := range groups {
//			list = append(list, v.Name)
//		}
//
//		Success(c, list)
//	}
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
		Fail(c, http.StatusOK, "绑定数据错误")
		return
	}
	name, err := dao.FilterUser(form.CN)
	if err != nil {
		Fail(c, http.StatusOK, err.Error())
		return
	}
	err = dao.AddUser(form.OU, name[0], form.CN, form.Pass)
	if err != nil {
		Fail(c, http.StatusOK, fmt.Sprintf("添加用户 %s失败:%s", name[0], err.Error()))
		return
	}
	err = AddUserForDB(form.OU, form.CN, name[0])
	if err != nil {
		Fail(c, http.StatusOK, fmt.Sprintf("用户%s信息入库失败: %s", name[0], err.Error()))
		return
	}
	Success(c, fmt.Sprintf("添加用户%s成功", name[0]))
}

func AddUserForDB(ou, cn, sn string) error {
	user := &model.User{
		Name:       cn,
		Email:      fmt.Sprintf("%s@lang.com", sn),
		DeptName:   ou,
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
		Fail(c, http.StatusOK, "绑定数据错误")
		return
	}

	err := dao.LdapDeleteUser(form.CN)
	if err != nil {
		Fail(c, http.StatusOK, fmt.Sprintf("删除用户%s失败: %s", form.CN, err))
		return
	}
	Success(c, fmt.Sprintf("删除用户%s成功", form.CN))
}

// 修改密码
func ResetPassword(c *gin.Context) {
	var form types.LdapResetPass
	if err := c.Bind(&form); err != nil {
		Fail(c, http.StatusOK, "绑定数据错误")
		return
	}
	sn, _ := dao.FindUserByLdap(form.CN)

	err := dao.LdapResetPassword(sn, form.Password)
	if err != nil {
		Fail(c, http.StatusOK, fmt.Sprintf("修改用户%s密码失败: %s", sn, err))
		return
	}
	Success(c, fmt.Sprintf("修改用户%s密码成功", sn))
}
