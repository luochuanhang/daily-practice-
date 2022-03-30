package util

import (
	"lianxi/107_blog/pkg/setting"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret = []byte(setting.JwtSecret)
//请求权
type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

//创建Token
func GenerateToken(username, password string) (string, error) {
	//当前时间
	nowTime := time.Now()
	//到期时间3小时
	expireTime := nowTime.Add(3 * time.Hour)

	//结构体初始化
	claims := Claims{
		username,
		password,
		//过期时间和发行人
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}
	//设置加密方案
	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//获取完整的签名令牌
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

//解析Token
func ParseToken(token string) (*Claims, error) {
	//解析验证并返回令牌
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	//token令牌不为空
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
