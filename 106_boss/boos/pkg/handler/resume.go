package handler

import (
	"lianxi/106_boss/boos/pkg/model"
	"lianxi/106_boss/boos/pkg/storage"
	"lianxi/106_boss/boos/pkg/validate"
	"lianxi/106_boss/boos/pkg/validate/regex"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type ResumeHandler struct {
	base
}

func (r ResumeHandler) Init(router *gin.RouterGroup, engine storage.Storage) {
	r.engine = engine
	r.logger = logrus.WithField("handler", "resume")
	resumegroup := router.Group("/resume")
	resumegroup.GET("/", r.QueryAll)
	resumegroup.GET("/:id", r.Query)
	resumegroup.POST("/", r.Add)
	resumegroup.PUT("/", r.Update)
	resumegroup.DELETE("/", r.Delete)
}
func (r ResumeHandler) Query(ctx *gin.Context) {
	reqs := []model.Resume{}
	phone := ctx.Param("id")
	if !validate.Regex(regex.Phone, []byte(phone)) {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "please checkout your Phone",
			"status":  "failed",
		})
	}
	db := r.engine.Get().Debug()
	result := db.Where("phone=?", phone).Find(&reqs)
	if result.Error != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "unknown error",
			"status":  "failed",
		})
		return
	}
	for _, req := range reqs {
		ctx.JSON(200, req)
	}
}

func (r ResumeHandler) QueryAll(ctx *gin.Context) {
	reqs := []model.Resume{}
	db := r.engine.Get().Debug()
	result := db.Find(&reqs)
	if result.Error != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "unknown error",
			"status":  "failed",
		})
		return
	}
	for _, req := range reqs {
		ctx.JSON(200, req)
	}

}
func (r ResumeHandler) Update(ctx *gin.Context) {
	req := model.Resume{}
	if err := ctx.Bind(&req); err != nil {
		r.badrequest(ctx)
		return
	}
	req.UpdateAt = time.Now()
	tx := r.engine.Get().Begin().Debug()
	result := tx.Model(&model.Resume{}).Where("phone=?", req.Phone).Updates(&req)
	if result.Error != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "please checkout your input",
			"status":  "Update failed",
		})
		tx.Rollback()
		return
	}
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "please checkout your input",
			"status":  "Update failed",
		})
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "Update success ",
	})
}
func (r ResumeHandler) Add(ctx *gin.Context) {
	req := model.Resume{}
	//参数绑定
	if err := ctx.Bind(&req); err != nil {
		r.badrequest(ctx)
		return
	}
	//检测手机号
	if !validate.Regex(regex.Phone, []byte(req.Phone)) {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "please checkout your Phone",
			"status":  "failed",
		})
		return
	}
	tx := r.engine.Get().Begin().Debug()
	//显示sql语句
	//增加一条记录
	result := tx.Create(&model.Resume{
		Phone:    req.Phone,
		Name:     req.Name,
		Age:      req.Age,
		Job:      req.Job,
		CreateAt: time.Now(),
		UpdateAt: time.Now(),
	})
	if result.Error != nil {
		r.logger.WithError(result.Error).Errorf("failed to regist")
		r.serverError(ctx)
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "Add success ",
	})

}
func (r ResumeHandler) Delete(ctx *gin.Context) {
	phone := ctx.Query("phone")
	tx := r.engine.Get().Begin().Debug()
	tx.Begin()
	result := tx.Where("phone=?", phone).Delete(&model.Resume{})
	if result.Error != nil {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "please checkout your input",
			"status":  "failed",
		})
		tx.Rollback()
		return
	}
	if result.RowsAffected == 0 {
		ctx.JSON(http.StatusOK, map[string]string{
			"message": "please checkout your input",
			"status":  "Delect failed",
		})
		tx.Rollback()
		return
	}
	tx.Commit()
	ctx.JSON(http.StatusOK, map[string]string{
		"status":  "ok",
		"message": "Delete success ",
	})
}
