package middleware

import (
	"github.com/gin-gonic/gin"
)

// GinMiddleware stands for gin middleware.
type GinMiddleware func() gin.HandlerFunc
