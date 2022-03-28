package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"lianxi/106_boss/boos/pkg/storage"
	"github.com/sirupsen/logrus"
)

func NewAuthenticate() GinMiddleware {
	return func() gin.HandlerFunc {
		return func(ctx *gin.Context) {
			if strings.Contains(ctx.Request.URL.Path, "login") {
				ctx.Next()
				return
			}
			if strings.Contains(ctx.Request.URL.Path, "regist") {
				ctx.Next()
				return
			}
			userId, err := ctx.Cookie("boos_session")
			if err != nil {
				ctx.AbortWithStatus(http.StatusUnauthorized)
				return
			}
			result := storage.Get().Get(fmt.Sprintf("session_%s", userId))
			logrus.Infof("get session from redis %s", result.String())
			if result.Err(); err != nil {
				ctx.AbortWithStatus(http.StatusInternalServerError)
				return
			}
		}
	}
}
