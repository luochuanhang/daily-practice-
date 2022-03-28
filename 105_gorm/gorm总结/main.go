package main

import (
	"fmt"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Student struct {
	Id         int
	Name       string
	Age        string
	CreateTime time.Time `gorm:"column:createtime"`
}

func (s Student) TableName() string {
	return "student"
}
func main() {
	db := GetDB()
	//表不存在就创建表
	if !db.Migrator().HasTable(&Student{}) {
		db.Migrator().CreateTable(&Student{})
	}
	s1 := Student{
		Name:       "luochuan",
		Age:        "5",
		CreateTime: time.Now(),
	}
	db.Where("id=?", 3).Debug().Updates(&s1)
	//db.Create(&s1)
	db.Where("id=2").Debug().Delete(&Student{})
	s2 := Student{}
	db.Take(&s2)
	fmt.Println(s2)
	s3 := []Student{}
	db.Find(&s3)
	fmt.Println(s3)
	// db.Transaction(func(tx *gorm.DB) error {
	// 	err := tx.Where("age=?", 5).Delete(&Student{}).Error
	// 	fmt.Println(err)
	// 	tx.Create(&s1)
	// 	return err
	// })
	tx := db.Begin()
	tx.Create(&s1)
	tx.Commit()
	tx = db.Begin()
	tx.Where("age=?", 5).Debug().Delete(&Student{})
	tx.Commit()
}

func GetDB() *gorm.DB {
	//配置MySQL连接参数
	username := "root"       //账号
	password := "123456"     //密码
	host := "172.18.100.186" //数据库地址，可以是Ip或者域名
	port := 3306             //数据库端口
	Dbname := "db01"         //数据库名
	str := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, port, Dbname)
	//建立数据库连接
	db, err := gorm.Open(mysql.Open(str))
	if err != nil {
		panic(err)
	}
	sqlDB, err := db.DB()
	sqlDB.SetMaxOpenConns(50) //设置连接池最大连接数
	sqlDB.SetMaxIdleConns(20) //连接池最大允许空连接数，没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
	return db
}
