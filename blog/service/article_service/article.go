package article_service

import (
	"encoding/json"

	"lianxi/blog/models"
	"lianxi/blog/pkg/gredis"
	"lianxi/blog/pkg/logging"
	"lianxi/blog/service/cache_service"
)

type Article struct {
	ID            int
	TagID         int
	Title         string
	Desc          string
	Content       string
	CoverImageUrl string
	State         int
	CreatedBy     string
	ModifiedBy    string

	PageNum  int
	PageSize int
}

//添加文章
func (a *Article) Add() error {
	article := map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"created_by":      a.CreatedBy,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
	}
	//添加文章到数据库
	if err := models.AddArticle(article); err != nil {
		return err
	}

	return nil
}

//编辑文章
func (a *Article) Edit() error {
	//将数据库中的数据更新
	return models.EditArticle(a.ID, map[string]interface{}{
		"tag_id":          a.TagID,
		"title":           a.Title,
		"desc":            a.Desc,
		"content":         a.Content,
		"cover_image_url": a.CoverImageUrl,
		"state":           a.State,
		"modified_by":     a.ModifiedBy,
	})
}

//获取单个文章
func (a *Article) Get() (*models.Article, error) {
	var cacheArticle *models.Article

	cache := cache_service.Article{ID: a.ID}
	//将key设置为文章+id
	key := cache.GetArticleKey()
	//检查是否存在
	if gredis.Exists(key) {
		//获取key中的数据
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticle)
			return cacheArticle, nil
		}
	}
	//根据id获取文章数据
	article, err := models.GetArticle(a.ID)
	if err != nil {
		return nil, err
	}
	//将对应的key和文章数据放入redis
	gredis.Set(key, article, 3600)
	return article, nil
}

//获取多个文章
func (a *Article) GetAll() ([]*models.Article, error) {
	var (
		articles, cacheArticles []*models.Article
	)
	//获取缓存
	cache := cache_service.Article{
		TagID: a.TagID,
		State: a.State,

		PageNum:  a.PageNum,
		PageSize: a.PageSize,
	}
	//获取key值
	key := cache.GetArticlesKey()
	//检查是否存在
	if gredis.Exists(key) {
		//获取redis中对应的数据
		data, err := gredis.Get(key)
		if err != nil {
			logging.Info(err)
		} else {
			json.Unmarshal(data, &cacheArticles)
			return cacheArticles, nil
		}
	}
	//获取多个数据
	articles, err := models.GetArticles(a.PageNum, a.PageSize, a.getMaps())
	if err != nil {
		return nil, err
	}
	//将数据放入redis中
	gredis.Set(key, articles, 3600)
	return articles, nil
}

//删除文章
func (a *Article) Delete() error {
	return models.DeleteArticle(a.ID)
}

//检查文章是否存在
func (a *Article) ExistByID() (bool, error) {
	return models.ExistArticleByID(a.ID)
}

//对文章进行计数
func (a *Article) Count() (int, error) {
	return models.GetArticleTotal(a.getMaps())
}

//未被删除,存在状态信息，存在tagid
func (a *Article) getMaps() map[string]interface{} {
	maps := make(map[string]interface{})
	maps["deleted_on"] = 0
	if a.State != -1 {
		maps["state"] = a.State
	}
	if a.TagID != -1 {
		maps["tag_id"] = a.TagID
	}

	return maps
}
