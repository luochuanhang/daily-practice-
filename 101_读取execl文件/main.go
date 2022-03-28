package main

import (
	"fmt"
	"sync"

	"github.com/xuri/excelize/v2"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Person struct {
	Field1 string `gorm:"column:field1"`
	Field2 string `gorm:"column:field2"`
	Field3 string `gorm:"column:field3"`
}

func (pe Person) TableName() string {
	//绑定MYSQL表名为resume
	return "person"
}

var (
	defaultDB *gorm.DB
	lock      = sync.Mutex{}
)

func NewDB() *gorm.DB {
	lock.Lock()
	defer lock.Unlock()
	if defaultDB == nil {
		//连接MYSQL, 获得DB类型实例
		dsn := "root:123456@tcp(172.18.108.99:3306)/person?charset=utf8&parseTime=True&loc=Local"
		db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			panic("数据库连接错误" + err.Error())
		}
		defaultDB = db
	}
	return defaultDB
}

func main() {
	f, err := excelize.OpenFile("表格.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	// 获取 Sheet1 上所有单元格
	rows, err := f.GetRows("Sheet1")
	if err != nil {
		fmt.Println(err)
		return
	}
	db := NewDB()
	for i, row := range rows {
		if i == 0 {
			// db.AutoMigrate(&Person{})
			continue
		}
		per := Person{}
		per.Field1 = row[0]
		per.Field2 = row[1]
		per.Field3 = row[2]
		//添加数据
		if err := db.Create(&per).Error; err != nil {
			fmt.Println("插入失败", err)
			return
		}
	}
}
