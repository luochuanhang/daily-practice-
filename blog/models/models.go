package models

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"lianxi/blog/pkg/setting"
	"time"
)

var db *gorm.DB

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedOn  int `json:"created_on"`
	ModifiedOn int `json:"modified_on"`
	DeletedOn  int `json:"deleted_on"`
}

// 设置数据库初始化实例
func Setup() {
	var err error
	//连接数据库
	db, err = gorm.Open(setting.DatabaseSetting.Type, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		setting.DatabaseSetting.User,
		setting.DatabaseSetting.Password,
		setting.DatabaseSetting.Host,
		setting.DatabaseSetting.Name))

	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}
	//默认的表名处理器
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return setting.DatabaseSetting.TablePrefix + defaultTableName
	}
	//默认使用单数表 表名不带s
	db.SingularTable(true)
	//
	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
	db.Callback().Delete().Replace("gorm:delete", deleteCallback)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

// 关闭数据库连接(unnecessary)
func CloseDB() {
	defer db.Close()
}

// 更新时间戳创建回调
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		//获取现在的时间
		nowTime := time.Now().Unix()
		//获取字段的数据
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			//检查字段是否是空的
			if createTimeField.IsBlank {
				//设置字段的值
				createTimeField.Set(nowTime)
			}
		}
		//获取字段的数据
		if modifyTimeField, ok := scope.FieldByName("ModifiedOn"); ok {
			//检查字段是否是空的
			if modifyTimeField.IsBlank {
				//设置字段的值
				modifyTimeField.Set(nowTime)
			}
		}
	}
}

// 更新时间戳更新回调
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	//按名称获取设置
	if _, ok := scope.Get("gorm:update_column"); !ok {
		//设置字段的值为当前时间
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}

// 删除回调
func deleteCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		var extraOption string
		//按名称获取设置
		if str, ok := scope.Get("gorm:delete_option"); ok {
			extraOption = fmt.Sprint(str)
		}
		//检查数据
		deletedOnField, hasDeletedOnField := scope.FieldByName("DeletedOn")
		//
		if !scope.Search.Unscoped && hasDeletedOnField {
			scope.Raw(fmt.Sprintf(
				"UPDATE %v SET %v=%v%v%v",
				scope.QuotedTableName(),
				scope.Quote(deletedOnField.DBName),
				scope.AddToVars(time.Now().Unix()),
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

// addExtraSpaceIfExist adds a separator
func addExtraSpaceIfExist(str string) string {
	if str != "" {
		return " " + str
	}
	return ""
}
