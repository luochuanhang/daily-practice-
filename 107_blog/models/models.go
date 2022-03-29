package models

import (
	"fmt"
	"lianxi/107_blog/pkg/setting"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
}

func init() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	//加载配置文件
	sec, err := setting.Cfg.GetSection("database")
	if err != nil {
		log.Fatal(2, "Fail to get section 'database': %v", err)
	}
	dbType = sec.Key("TYPE").String()
	dbName = sec.Key("NAME").String()
	user = sec.Key("USER").String()
	password = sec.Key("PASSWORD").String()
	host = sec.Key("HOST").String()
	tablePrefix = sec.Key("TABLE_PREFIX").String()
	//连接数据库
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))

	if err != nil {
		log.Println(err)
	}
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		//返回表前缀+默认表名
		return tablePrefix + defaultTableName
	}
	//SingularTable默认使用单数表
	db.SingularTable(true)
	//LogMode设置日志模式，'true'表示详细日志，
	//'false'表示无日志，默认情况下，将只打印错误日志
	db.LogMode(true)
	//设置连接池最大连接100，最大空闲10
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}
//关闭连接
func CloseDB() {
	defer db.Close()
}
