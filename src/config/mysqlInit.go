package config

import (
	"fmt"
	util "testing/src/util"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func GetDB() (db *gorm.DB) {
	if db == nil {
		InitDB()
	}
	return db
}

func InitDB() {
	var username string = util.GetElement("db.user")
	var password string = util.GetElement("db.password")
	var host string = util.GetElement("db.host")
	var port int = util.StrToInt(util.GetElement("db.port"))
	var dbname string = util.GetElement("db.dbname")
	var timeout string = util.GetElement("db.timeouts")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local&timeout=%s", username, password, host, port, dbname, timeout)
	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("connected to db fail, error=" + err.Error())
	}
	fmt.Println("db init......")
	sqlDB, _ := db.DB()
	sqlDB.SetMaxOpenConns(util.StrToInt(util.GetElement("db.maxconnections")))
	sqlDB.SetMaxIdleConns(util.StrToInt(util.GetElement("db.maxidleconns")))

}
