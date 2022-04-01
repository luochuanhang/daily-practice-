package models

import (
	"fmt"
	"lianxi/107_blog/pkg/setting"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

//初始化数据库连接
func Setup() {
	var (
		err                                               error
		dbType, dbName, user, password, host, tablePrefix string
	)
	//加载配置文件

	dbType = setting.DatabaseSetting.Type
	dbName = setting.DatabaseSetting.Name
	user = setting.DatabaseSetting.User
	password = setting.DatabaseSetting.Password
	host = setting.DatabaseSetting.Host
	tablePrefix = setting.DatabaseSetting.TablePrefix
	//连接数据库
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, dbName))

	if err != nil {
		log.Println(err)
	}
	//默认表名处理器
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
	//回调函数增加替换
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
}

//创建更新时间回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	//如果没有错误
	if !scope.HasError() {
		//现在的时间
		nowTime := time.Now().Unix()
		//通过scope.Fields()获取所有字段，判断当前是否包含所需字段
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			//field.IsBlank 可判断该字段的值是否为空
			if createTimeField.IsBlank {
				//若为空则 field.Set 用于给该字段设置值
				createTimeField.Set(nowTime)
			}
		}

		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			if modifyTimeField.IsBlank {
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

//更新时回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//根据入参获取设置了字面值的参数， gorm:update_column ，它会去查找含这个字面值的字段属性
	if _, ok := scope.Get("gorm:update_column"); !ok {
		//假设没有指定 update_column 的字段，我们默认在更新回调设置 ModifiedOn 的值
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

//删除时回调
func deleteCallback(scope *gorm.Scope) {
	//如果没有错误
	if !scope.HasError() {
		var extraOption string
		//检查是否手动指定了 delete_option
		if str, ok := scope.Get("gorm:delete_option"); ok {
			//有数据则设置变量数据
			extraOption = fmt.Sprint(str)
		}
		//获取我们约定的删除字段，若存在则UPDATE软删除，若不存在则DELETE硬删除
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		if !scope.Search.Unscoped && hasDeletedOnField {

			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				//返回引用的表名，这个方法 GORM 会根据自身逻辑对表名进行一些处理
				scope.QuotedTableName(),
				//用于引用字符串来为数据库转义它们
				scope.Quote(deletedOnField.DBName),
				//添加值作为sql的变量，用于防止sql注入
				scope.AddToVars(time.Now().Unix()),
				//CombinedConditionSql返回组合条件sql
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		} else {
			scope.Raw(fmt.Sprintf(
				"DELETE FROM %v%v%v",
				scope.QuotedTableName(),
				addExtraSpaceIfExist(scope.CombinedConditionSql()),
				addExtraSpaceIfExist(extraOption),
			)).Exec()
		}
	}
}

func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}

//关闭连接
func CloseDB() {
	defer db.Close()
}
