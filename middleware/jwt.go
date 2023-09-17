package middleware

import (
	"github.com/Sion-L/onePiece/types"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var jwtSecret = []byte("one-piece")

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	//Role     string `json:"role"`
	jwt.StandardClaims
}

func ReleaseToken(userName string, password string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserName: userName,
		Password: password,
		//Role:     role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			IssuedAt:  time.Now().Unix(), // 发放时间
			Issuer:    "lang",            // 发放人
			Subject:   "user token",      // 主题
		}}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func ParseToken(token string) (*Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}

func JWTMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("One-Piece")
		if tokenString == "" {
			types.Fail(c, http.StatusUnauthorized, "Token not active yet")
			c.Abort()
			return
		} else {
			claims, err := ParseToken(tokenString)
			if err != nil {
				types.FailF(c, http.StatusUnauthorized, "Failed to parse token", err.Error())
			} else if time.Now().Unix() > claims.ExpiresAt {
				types.Fail(c, http.StatusUnauthorized, "Token is expired")
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

// TODO:
// 1.后端登陆校验完之后根据前端传过来的参数生成token添加到响应头中 完成
// 2.中间件校验前端传过来的token 完成
// 3.token过期后进行刷新(暂不需要)

// 前端从响应头获取token并存入请求头中，每次请求都携带token发送
