package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lianxi/blog/pkg/app"
	"lianxi/blog/pkg/e"
	"lianxi/blog/pkg/logging"
	"lianxi/blog/pkg/upload"
)

// @Summary Import Image
// @Produce  json
// @Param image formData file true "Image File"
// @Success 200 {object} app.Response
// @Failure 500 {object} app.Response
// @Router /api/v1/tags/import [post]
func UploadImage(c *gin.Context) {
	appG := app.Gin{C: c}
	//获取form表单文件
	file, image, err := c.Request.FormFile("image")
	if err != nil {
		//写入log文件
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR, nil)
		return
	}

	if image == nil {
		//返回json数据
		appG.Response(http.StatusBadRequest, e.INVALID_PARAMS, nil)
		return
	}
	//获取图片的名字
	imageName := upload.GetImageName(image.Filename)
	//获取完整路径
	fullPath := upload.GetImageFullPath()
	//获取保存地址
	savePath := upload.GetImagePath()
	//完整的路径+文件名称
	src := fullPath + imageName
	//检查文件的后缀和大小是否满足要求
	if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
		appG.Response(http.StatusBadRequest, e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT, nil)
		return
	}
	//检查文件地址是否存在
	err = upload.CheckImage(fullPath)
	if err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_CHECK_IMAGE_FAIL, nil)
		return
	}
	//SaveUploadedFile将表单文件上传到指定的地址
	if err := c.SaveUploadedFile(image, src); err != nil {
		logging.Warn(err)
		appG.Response(http.StatusInternalServerError, e.ERROR_UPLOAD_SAVE_IMAGE_FAIL, nil)
		return
	}
	//返回文件地址和保存地址
	appG.Response(http.StatusOK, e.SUCCESS, map[string]string{
		"image_url":      upload.GetImageFullUrl(imageName),
		"image_save_url": savePath + imageName,
	})
}
