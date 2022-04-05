package util

import (
	"time"

	"github.com/dgrijalva/jwt-go"
)

var jwtSecret []byte

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	jwt.StandardClaims
}

// GenerateToken生成用于认证的令牌
func GenerateToken(username, password string) (string, error) {
	//当前时间
	nowTime := time.Now()
	//过期时间3个小时   当前时间+3小时
	expireTime := nowTime.Add(3 * time.Hour)
	//创建claims实例
	claims := Claims{
		EncodeMD5(username),
		EncodeMD5(password),
		//jwt过期时间和发行人
		jwt.StandardClaims{
			ExpiresAt: expireTime.Unix(),
			Issuer:    "gin-blog",
		},
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	//获取完整的签名令牌
	token, err := tokenClaims.SignedString(jwtSecret)
	return token, err
}

// ParseToken解析令牌
func ParseToken(token string) (*Claims, error) {
	/*
		警告:除非你知道你在做什么，否则不要使用这个方法此方法解析令牌，但不验证签名。
		它只在您知道签名是有效的(因为它以前在堆栈中检查过)并且您想从中提取值的情况下才有用。
	*/
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})

	if tokenClaims != nil {
		//tokenClaims.Claims令牌的第二段//tokenClaims.Valid检查令牌是否有效，在解析/验证令牌时填充
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}

	return nil, err
}
