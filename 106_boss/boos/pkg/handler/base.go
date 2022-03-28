package handler

import (
	"net/http"

	"lianxi/106_boss/boos/pkg/storage"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type base struct {
	engine storage.Storage
	logger *logrus.Entry
}

func (b base) badrequest(ctx *gin.Context) {
	ctx.JSON(http.StatusBadRequest, map[string]interface{}{
		"message": "please checkout your input",
		"status":  "failed",
	})
}

func (b base) serverError(ctx *gin.Context) {
	ctx.JSON(http.StatusInternalServerError, map[string]string{
		"message": "please checkout your input",
		"status":  "failed",
	})
}

func (b base) ok(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, map[string]string{
		"message": "success",
		"status":  "success",
	})
}
