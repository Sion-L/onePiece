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
	"github.com/Sion-L/onePiece/dao"
	"github.com/gin-gonic/gin"
)

func AddUserForLdap(c *gin.Context) {

	dao.AddUser()
}
