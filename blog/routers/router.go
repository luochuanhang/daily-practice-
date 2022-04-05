package routers

import (
	"net/http"

	"github.com/gin-gonic/gin"

	_ "lianxi/blog/docs"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"

	"lianxi/blog/middleware/jwt"
	"lianxi/blog/pkg/export"
	"lianxi/blog/pkg/qrcode"
	"lianxi/blog/pkg/upload"
	"lianxi/blog/routers/api"
	v1 "lianxi/blog/routers/api/v1"
)

// InitRouter 初始化路由信息
func InitRouter() *gin.Engine {
	//New返回一个新的空引擎实例，不附带任何中间件。
	r := gin.New()
	//添加中间件
	//Logger实例一个Logger中间件，该中间件将日志写入gin.DefaultWriter。
	r.Use(gin.Logger())
	//Recovery返回一个中间件，它可以从任何恐慌中恢复过来，如果有的话，它会写一个500。
	r.Use(gin.Recovery())
	/*
	StaticFS的工作原理就像'Static()'，但一个自定义的'http。
	可以使用FileSystem'来代替。Gin默认用户:Gin . dir ()
	*/
	//Dir使用限制在特定目录树中的本机文件系统实现文件系统。
	r.StaticFS("/export", http.Dir(export.GetExcelFullPath()))
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))
	r.StaticFS("/qrcode", http.Dir(qrcode.GetQrCodeFullPath()))

	r.POST("/auth", api.GetAuth)
	//swagger调试
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//上传图片
	r.POST("/upload", api.UploadImage)
	//创建一个路由组
	apiv1 := r.Group("/api/v1")
	//添加一个jwt中间件进行token验证
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("/tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("/tags/:id", v1.DeleteTag)
		//导出标签
		r.POST("/tags/export", v1.ExportTag)
		//导入标签
		r.POST("/tags/import", v1.ImportTag)

		//获取文章列表
		apiv1.GET("/articles", v1.GetArticles)
		//获取指定文章
		apiv1.GET("/articles/:id", v1.GetArticle)
		//新建文章
		apiv1.POST("/articles", v1.AddArticle)
		//更新指定文章
		apiv1.PUT("/articles/:id", v1.EditArticle)
		//删除指定文章
		apiv1.DELETE("/articles/:id", v1.DeleteArticle)
		//生成文章海报
		apiv1.POST("/articles/poster/generate", v1.GenerateArticlePoster)
	}

	return r
}
