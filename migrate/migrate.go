package migrate

import (
	"fmt"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/jinzhu/gorm"
	"github.com/asdfsx/zhihu-golang-web/models"
)

var db *gorm.DB
var modelsRegisted []interface{} = []interface{}{&models.Auth{}, &models.Product{}, &models.Category{}}

func connect(dbType, user, password, host string, port int64, dbName, tablePrefix string) (err error) {
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	if err != nil{
		jww.ERROR.Println(err)
		return err
	}
	return nil
}

func createDatabaseIfNotExists(dbType, user, password, host string, port int64, dbName, tablePrefix string) (err error) {
	err = connect(dbType, user, password, host , port , "mysql", tablePrefix)
	if err != nil{
		jww.ERROR.Println(err)
		return err
	}
	defer db.Close()

	err = db.Exec(fmt.Sprintf("create database if not exists `%s`", dbName)).Error
	if err != nil{
		jww.ERROR.Println(err)
		return err
	}
	return nil
}

func ConnectDB(dbType, user, password, host string, port int64, dbName, tablePrefix string) (err error) {
	err = createDatabaseIfNotExists(dbType, user, password, host , port , dbName, tablePrefix )
	if err != nil{
		jww.ERROR.Println(err)
		return err
	}

	err = connect(dbType, user, password, host , port , dbName, tablePrefix)
	if err != nil{
		jww.ERROR.Println(err)
		return err
	}
	db.LogMode(true)
	db.SingularTable(true)
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName;
	}
	return err
}

func Close() error {
	return db.Close()
}

func Migrate() error {
	if err := db.AutoMigrate(modelsRegisted).Error;err != nil{
		jww.ERROR.Println(err)
		return err
	}
	return nil
}