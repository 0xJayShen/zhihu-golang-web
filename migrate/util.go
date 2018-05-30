package migrate

import (
	"fmt"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

func CreateDBConnection(dbType, user, password, host string, port int64, dbName string) (err error) {
	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		user,
		password,
		host,
		dbName))
	gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
		return tablePrefix + defaultTableName;
	}
	db.LogMode(true)
	db.SingularTable(true)
	return err
}

func Close() error {
	return db.Close()
}
