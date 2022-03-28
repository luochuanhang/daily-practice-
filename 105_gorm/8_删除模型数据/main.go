package main

import tools "lianxi/105_gorm/4_tools"

type Food struct {
	Id int
	//常用标签 columb 列名
	//指定主键 primarykey
	//- 忽略字段
	Name       string
	Price      float64
	TypeId     int
	CreateTime int64 `gorm:"column:createtime"`
}

func (f Food) TableName() string {
	return "food"
}
func main() {
	db := tools.GetDB()
	food := Food{}
	//先查询一条记录保存到模型
	db.Where("id=?", 2).Take(&food)
	//根据模型主键删除数据
	db.Debug().Delete(&food)
	//根据条件删除
	db.Where("id=?", 1).Delete(&Food{})
}
