package tools

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

//每个包被调用时会优先执行init 初始化
func init() {
	var err error
	username := "root"
	password := "123456"
	host := "172.18.100.186"
	port := 3306
	Dbname := "db01"
	timeout := "10s"
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, Dbname, timeout)
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("连接错误")
	}
	sqlDB, _ := db.DB()
	//设置连接池参数
	sqlDB.SetMaxOpenConns(50) //设置连接池最大连接数
	sqlDB.SetMaxIdleConns(20) //连接池最大允许空连接数，没有sql任务需要执行的连接数大于20，超过的连接会被连接池关闭。
}

//调用getdb 返回db
func GetDB() *gorm.DB {
	return db
}
