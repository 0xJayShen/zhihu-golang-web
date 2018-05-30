package models

import (

	"github.com/jinzhu/gorm"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"time"
	//"gin-docker-mysql/pkg/logging"
)

var db *gorm.DB


//func init() {
//	var (
//		err                                               error
//		dbType, dbName, user, password, host, tablePrefix string
//	)
//
//	dbType = setting.DataBase_.TYPE
//	dbName = setting.DataBase_.NAME
//	user = setting.DataBase_.USER
//	password = setting.DataBase_.PASSWORD
//	host = setting.DataBase_.HOST
//	tablePrefix = setting.DataBase_.TABLE_PREFIX
//	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
//		user,
//		password,
//		host,
//		dbName))
//
//	if err != nil {
//		fmt.Println(err, "-------")
//	}
//	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
//		return tablePrefix + defaultTableName;
//	}
//	db.LogMode(true)
//	db.SingularTable(true)
//	//db.AutoMigrate(&Product{})
//
//	db.Callback().Create().Replace("gorm:update_time_stamp", updateTimeStampForCreateCallback)
//	db.Callback().Update().Replace("gorm:update_time_stamp", updateTimeStampForUpdateCallback)
//	db.DB().SetMaxIdleConns(10)
//	db.DB().SetMaxOpenConns(100)
//}

func CloseDB() {
	defer db.Close()
}

// updateTimeStampForCreateCallback will set `CreatedOn`, `ModifiedOn` when creating
func updateTimeStampForCreateCallback(scope *gorm.Scope) {
	if !scope.HasError() {
		nowTime := time.Now().Unix()
		if createTimeField, ok := scope.FieldByName("CreatedOn"); ok {
			if createTimeField.IsBlank {
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

// updateTimeStampForUpdateCallback will set `ModifiedOn` when updating
func updateTimeStampForUpdateCallback(scope *gorm.Scope) {
	if _, ok := scope.Get("gorm:update_column"); !ok {
		scope.SetColumn("ModifiedOn", time.Now().Unix())
	}
}
