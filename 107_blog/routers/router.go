package routers

import (
	_ "lianxi/107_blog/docs"
	"lianxi/107_blog/middleware/jwt"
	"lianxi/107_blog/pkg/upload"
	"lianxi/107_blog/routers/api"
	v1 "lianxi/107_blog/routers/api/v1"
	"net/http"

	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode("debug")
	//StaticFS的工作原理就像'Static()'，但一个自定义的'http。可以使用FileSystem'来代替。
	/*
		Dir使用限制在特定目录的本机文件系统实现文件系统树。
		当文件系统。Open方法接受“分开的路径，一个Dir的字符串
		值是本地文件系统上的文件名，而不是URL，
		所以它被filepath分隔。”分离器,
	*/
	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	r.GET("/auth", api.GetAuth)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.POST("/upload", api.UploadImage)
	apiv1 := r.Group("/api/v1")
	apiv1.Use(jwt.JWT())
	{
		//获取标签列表
		apiv1.GET("/tags", v1.GetTags)
		//新建标签
		apiv1.POST("/tags", v1.AddTag)
		//更新指定标签
		apiv1.PUT("tags/:id", v1.EditTag)
		//删除指定标签
		apiv1.DELETE("tags/:id", v1.DeleteTag)
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
	}

	return r
}
