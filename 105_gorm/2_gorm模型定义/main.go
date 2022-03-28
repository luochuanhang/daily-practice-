package main

import (
	"time"

	"gorm.io/gorm"
)

//gorm定义了一个结构体
type Model struct {
	ID        uint `gorm:"primaryKey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
}
type Food struct {
	//常用标签 columb 列名
	//指定主键 primarykey
	//- 忽略字段
	gorm.Model //嵌入gorm.Model 类似继承
	Name       string
	Price      float64
	TypeId     int
	CreateTime int64 `gorm:"column:createtime"`
}

//设置表名可以通过TableName函数返回表名
func (f Food) TableName() string {
	return "food"
}
func main() {

}
