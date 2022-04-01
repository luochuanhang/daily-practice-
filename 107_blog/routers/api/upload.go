package api

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"lianxi/107_blog/pkg/logging"
	"lianxi/107_blog/pkg/setting/e"
	"lianxi/107_blog/pkg/upload"
)

//上传图片
func UploadImage(c *gin.Context) {
	//错误码成功
	code := e.SUCCESS
	//存放数据
	data := make(map[string]string)
	//获取form图片文件
	file, image, err := c.Request.FormFile("image")
	//如果有错
	if err != nil {
		//写入日志
		logging.Warn(err)
		//修改错误码
		code = e.ERROR
		//返回错误
		c.JSON(http.StatusOK, gin.H{
			"code": code,
			"msg":  e.GetMsg(code),
			"data": data,
		})
	}
	//如果没有
	if image == nil {
		code = e.INVALID_PARAMS
	} else {
		imageName := upload.GetImageName(image.Filename)
		fullPath := upload.GetImageFullPath()
		savePath := upload.GetImagePath()

		src := fullPath + imageName
		//检查文件的后缀和大小
		if !upload.CheckImageExt(imageName) || !upload.CheckImageSize(file) {
			//有错设置错误码
			code = e.ERROR_UPLOAD_CHECK_IMAGE_FORMAT
		} else {
			//检查文件目录是否存在能不能访问
			err := upload.CheckImage(fullPath)
			if err != nil {
				//有错打印日志返回错误码
				logging.Warn(err)
				code = e.ERROR_UPLOAD_CHECK_IMAGE_FAIL
				//SaveUploadedFile将表单文件上传到指定的dst。
			} else if err := c.SaveUploadedFile(image, src); err != nil {
				logging.Warn(err)
				code = e.ERROR_UPLOAD_SAVE_IMAGE_FAIL
			} else {
				//给map放入文件的信息
				data["image_url"] = upload.GetImageFullUrl(imageName)
				data["image_save_url"] = savePath + imageName
			}
		}
	}
	//返回给前端
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
