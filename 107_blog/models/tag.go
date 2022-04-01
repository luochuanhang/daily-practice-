package models

import (
	"time"

	"github.com/jinzhu/gorm"
)

//编辑标签列表的models
type Tag struct {
	Model

	Name       string `json:"name"`
	CreatedBy  string `json:"created_by"`
	ModifiedBy string `json:"modified_by"`
	State      int    `json:"state"`
}

//获取多个文章标签
func GetTags(pageNum int, pageSize int, maps interface{}) (tags []Tag) {
	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&tags)

	return
}

//获取标签的总数
func GetTagTotal(maps interface{}) (count int) {
	db.Model(&Tag{}).Where(maps).Count(&count)

	return
}

//
func ExitTagByName(name string) bool {
	var tag Tag
	db.Select("id").Where("name=?", name).First(&tag)
	return tag.ID > 0
}

//新增标签
func AddTag(name string, state int, createBy string) bool {
	db.Create(&Tag{
		Name:      name,
		State:     state,
		CreatedBy: createBy,
	})

	return true
}

func (tag *Tag) BeforeCreate(scope *gorm.Scope) error {
	scope.SetColumn("createdOn", time.Now().Unix())
	return nil
}
func (tag *Tag) BeforeUpdate(scope *gorm.Scope) error {
	scope.SetColumn("ModifiedOn", time.Now().Unix())
	return nil
}
func ExistTagByID(id int) bool {
	var tag Tag
	db.Select("id").Where("id = ?", id).First(&tag)
	return tag.ID > 0
}
func DeleteTag(id int) bool {
	db.Where("id = ?", id).Delete(&Tag{})

	return true
}

func EditTag(id int, data interface{}) bool {

	db.Model(&Tag{}).Where("id = ?", id).Updates(data)

	return true
}
func CleanAllTag() bool {
	//注意硬删除要使用 Unscoped()，这是 GORM 的约定
	db.Unscoped().Where("deleted_on != ? ", 0).Delete(&Tag{})

	return true
}
