package handler

import (
	"lianxi/106_boss/boos/pkg/model"
	"lianxi/106_boss/boos/pkg/storage"
	"lianxi/106_boss/boos/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type RegistHandler struct {
	base
}

func (r *RegistHandler) Init(router *gin.RouterGroup, engine storage.Storage) {
	r.engine = engine
	r.logger = logrus.WithField("handler", "regist")
	router.POST("/regist", r.Regist)
}

func (r *RegistHandler) Regist(ctx *gin.Context) {
	var req model.Regist
	if err := ctx.Bind(&req); err != nil {
		r.badrequest(ctx)
		return
	}
	db := r.engine.Get().Debug()
	result := db.Create(&model.User{
		Username: req.Username,
		Password: util.Password(req.Password),
	})
	if result.Error != nil {
		r.logger.WithError(result.Error).Errorf("failed to regist")
		r.serverError(ctx)
		return
	}
	r.ok(ctx)
}
