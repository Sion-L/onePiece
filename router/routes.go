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
	auth := eng.Group("/api/v1/auth")

	// 添加jwt
	//auth.Use(middleware.JWTMiddleware())

	// 登陆
	auth.POST("login", controller.Login)

	// 添加用户
	auth.POST("addUser", controller.AddUserForLdap)

	// 删除用户
	auth.POST("deleteUser", controller.DeleteUser)

	// 修改密码
	auth.POST("changePassword", controller.ResetPassword)

	// 查询所有用户信息
	auth.GET("getAllUser", controller.FindAllUser)
}
