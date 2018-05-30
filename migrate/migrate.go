package migrate

import (
	"fmt"
	jww "github.com/spf13/jwalterweatherman"
	"github.com/jinzhu/gorm"
	"github.com/asdfsx/zhihu-golang-web/models"
)

var db *gorm.DB
var modelsRegisted []interface{} = []interface{}{&models.Auth{}, &models.Product{}, &models.Category{}}

func CreateDBConnection(dbType, user, password, host string, port int64, dbName, tablePrefix string) (err error) {
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName;
	}
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	db.LogMode(true)
	db.SingularTable(true)
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