package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Article struct {
	Model

	TagID int `json:"tag_id" gorm:"index"`
	Tag   Tag `json:"tag"`

	Title      string `json:"title"`
	Desc       string `json:"desc"`
	Content    string `json:"content"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//新增数据发生错误将会回滚
func (article *Article) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("CreatedOn", time.Now().Unix())

	return nil
}

//更新数据发生错误将会回滚
func (article *Article) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())

	return nil
}

//是否存在这个id
func ExistArticleByID(id int) bool {
	var article Article
	//查询表中的id将第一个赋值给文章模型
	db.Select("id").Where("id = ?", id).First(&article)
	//如果有数据则赋值给article 大于0就代表有数据
	return article.ID > 0

}

//查询文章的总数
func GetArticleTotal(maps interface{}) (count int) {
	db.Model(&Article{}).Where(maps).Count(&count)

	return
}

func GetArticles(pageNum int, pageSize int, maps interface{}) (articles []Article) {
	db.Preload("Tag").Where(maps).Offset(pageNum).Limit(pageSize).Find(&articles)

	return
}

//查询id对应的文章，和相关的文章标签
func GetArticle(id int) (article Article) {
	db.Where("id = ?", id).First(&article)
	db.Model(&article).Related(&article.Tag)

	return
}

//更新文章
func EditArticle(id int, data interface{}) bool {
	db.Model(&Article{}).Where("id = ?", id).Updates(data)

	return true
}

//添加文章
func AddArticle(data map[string]interface{}) bool {
	db.Create(&Article{
		TagID:     data["tag_id"].(int),
		Title:     data["title"].(string),
		Desc:      data["desc"].(string),
		Content:   data["content"].(string),
		CreatedBy: data["created_by"].(string),
		State:     data["state"].(int),
	})

	return true
}

//删除文章
func DeleteArticle(id int) bool {
	db.Where("id = ?", id).Delete(Article{})

	return true
}

func CleanAllArticle() bool {
	//注意硬删除要使用 Unscoped()，这是 GORM 的约定
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Article{})
	return true
}
