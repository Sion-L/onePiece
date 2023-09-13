package types

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

//func Response(c *gin.Context, httpStatus, code int, data gin.H, msg string) {
//	c.JSON(httpStatus, gin.H{
//		"code": code,
//		"data": data,
//		"msg":  msg,
//	})
//}
//
//// gin包装成功日志 - json
//func Success(c *gin.Context, data gin.H, msg string) {
//	Response(c, http.StatusOK, 200, data, msg)
//}
//
//// gin包装失败日志 - json
//func Fail(c *gin.Context, msg string, data gin.H) {
//	Response(c, http.StatusOK, 400, data, msg)
//}

// gin包装成功日志 - json
func Success(c *gin.Context, data interface{}) {
	c.JSON(http.StatusOK, gin.H{
		"code": 200,
		"msg":  "success",
		"data": data,
	})
}

// gin包装失败日志 - json
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": nil,
	})
}

func FailF(c *gin.Context, code int, msg string, err error) {
	c.JSON(http.StatusOK, gin.H{
		"code":   code,
		"msg":    msg,
		"data":   nil,
		"detail": err,
	})
}
