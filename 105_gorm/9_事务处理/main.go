package main

import (
	"fmt"
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
	db := tools.GetDB()
	//transation  自动事务 有问题就回滚
	db.Transaction(func(tx *gorm.DB) error {
		// 在事务中执行一些 db 操作（从这里开始，您应该使用 'tx' 而不是 'db'）
		if err := tx.Create(&Food{Name: "Giraffe"}).Error; err != nil {
			// 返回任何错误都会回滚事务
			return err
		}

		if err := tx.Create(&Food{Name: "Lion"}).Error; err != nil {
			return err
		}

		// 返回 nil 提交事务
		return nil
	})
	//手动事务
	//begin开始  开启事务
	tx := db.Begin()
	i := tx.Model(&Food{}).Where("id=?", 4).Update("name", "wangwu").RowsAffected
	fmt.Println(i)
	if i == 0 {
		//没有修改数据就回滚
		tx.Rollback()
		return
	}
	//没问题就提交
	tx.Commit()
}
