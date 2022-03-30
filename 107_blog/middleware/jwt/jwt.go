package jwt

import (
	"lianxi/107_blog/pkg/setting/util"
	"net/http"
	"time"

	"lianxi/107_blog/pkg/setting/e"

	"github.com/gin-gonic/gin"
)

//JWT中间件 是一个handler
func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		//错误码
		var code int
		//数据
		var data interface{}
		//错误码正常
		code = e.SUCCESS
		//获取URL中的token字段
		token := c.Query("token")
		//如果为空
		if token == "" {
			//错误码设置有错
			code = e.INVALID_PARAMS
		} else {
			//解析token
			claims, err := util.ParseToken(token)
			//如果没报错
			if err != nil {
				//错误码设置对应的错误
				code = e.ERROR_AUTH_CHECK_TOKEN_FAIL
				//如果现在的时间大于过期的时间
			} else if time.Now().Unix() > claims.ExpiresAt {
				//返回对应的错误码
				code = e.ERROR_AUTH_CHECK_TOKEN_TIMEOUT
			}
		}
		//如果错误码不是成功
		if code != e.SUCCESS {
			//返回http401，打印错误信息
			c.JSON(http.StatusUnauthorized, gin.H{
				"code": code,
				"msg":  e.GetMsg(code),
				"data": data,
			})
			//不执行后面的内容
			c.Abort()
			return
		}
		//执行后面的内容
		c.Next()
	}
}
