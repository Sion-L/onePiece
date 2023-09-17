package main

import (
	"github.com/Sion-L/onePiece/db"
	"github.com/Sion-L/onePiece/middleware"
	"github.com/Sion-L/onePiece/router"
	"github.com/gin-gonic/gin"
	"time"
)

func init() {
	db.InitLdap()
	db.InitMySQLDB()
}

func main() {

	r := gin.New()
	logger := middleware.GetLogger()
	r.Use(middleware.GinZap(logger, time.RFC3339, true))

	// Logs all panic to error log
	//   - stack means whether output the stack info.
	r.Use(middleware.RecoveryWithZap(logger, true))
	// 注册路由
	router.InitApi(r)

	// 启动服务
	if err := r.Run(":8888"); err != nil {
		logger.Debug("warn")
	}
}
