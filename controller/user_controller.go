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
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

type LdapUser struct {
	OU string `json:"ou" binding:"required"`
	CN string `json:"cn" binding:"required"`
	//SN   string   输入的是中文名,其内转换即可
	// UID  string	`json:"uid"` 等同于sn
	// Mail string  sn后加@xxx.xxx
	Pass string `json:"pass" binding:"required"`
}

func AddUserForLdap(c *gin.Context) {
	var form LdapUser
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
	var res struct {
		Cn string `json:"cn"`
	}
	if err := c.Bind(&res); err != nil {
		Fail(c, http.StatusOK, "绑定数据错误")
		return
	}

	err := dao.LdapDeleteUser(res.Cn)
	if err != nil {
		Fail(c, http.StatusOK, fmt.Sprintf("删除用户%s失败: %s", res.Cn, err))
		return
	}
	Success(c, fmt.Sprintf("删除用户%s成功", res.Cn))
}
