package main

import (
	tools "lianxi/105_gorm/4_tools"

	"gorm.io/gorm"
)

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
	//save 用于保存模型变量的值
	//根据主键id，更新所有模型字段
	db := tools.GetDB()
	f := Food{}
	//先查询一条记录，保存在模型变量
	db.Where("id=?", 2).Take(&f)
	//修改模型的值
	f.Price = 99
	db.Debug().Save(&f)
	// Update 更新单个字段值
	//修改模型的id中某个值
	db.Model(&f).Debug().Update("price", 69)
	//修改模型中的特定字段的值
	db.Model(&Food{}).Where("price>?", 50).Update("price", 88)

	//updates 更新多个字段值
	updaFood := Food{
		Name:  "hebei",
		Price: 95,
	}
	db.Model(&f).Updates(&updaFood)
	//根据条件更新多个数据
	db.Model(&Food{}).Where("price > ?", 30).Updates(&updaFood)
	data := make(map[string]interface{})
	data["name"] = "" //零值字段
	data["price"] = 35
	db.Model(&Food{}).Where("id = ?", 2).Updates(data)

	//更新表达式
	//gorm提供了Expr函数用于设置表达式
	db.Model(&f).Update("price", gorm.Expr("price + 1"))
}
