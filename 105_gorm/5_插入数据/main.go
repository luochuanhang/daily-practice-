package main

import (
	"fmt"
	tools "lianxi/105_gorm/4_tools"
	"time"
)

type Food struct {
	//常用标签 columb 列名
	//指定主键 primarykey
	//- 忽略字段
	Id         int
	Name       string
	Price      float64
	TypeId     int
	CreateTime int64 `gorm:"column:createtime"`
}

func (f Food) TableName() string {
	return "food"
}
func main() {
	f := Food{Name: "zhansan",
		Price:      45.5,
		TypeId:     5,
		CreateTime: time.Now().Unix()}
	db := tools.GetDB()
	if err := db.Debug().Create(&f).Error; err != nil {
		fmt.Println("插入失败")
		return
	}
	fmt.Println(f.Id)
}
