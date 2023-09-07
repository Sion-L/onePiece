package router

import (
	"github.com/Sion-L/onePiece/controller"
	"github.com/Sion-L/onePiece/middleware"
	"github.com/gin-gonic/gin"
)

func InitApi(eng *gin.Engine) {

	// 使用跨域中间件
	eng.Use(middleware.CoreMiddleware)

	// 接口分组
	api := eng.Group("/api/v1")

	// 添加用户
	api.POST("addUser", controller.AddUserForLdap)

	// 删除用户
	api.POST("deleteUser", controller.DeleteUser)

	// 修改密码
	api.POST("changePassword", controller.ResetPassword)

}
