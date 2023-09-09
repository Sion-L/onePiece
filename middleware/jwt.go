package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

var jwtSecret = []byte("one-piece")

type Claims struct {
	UserName string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

func ReleaseToken(userName string, password string) (string, error) {
	expireTime := time.Now().Add(24 * time.Hour)
	claims := &Claims{
		UserName: userName,
		Password: password,
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
