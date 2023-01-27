package common

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"simple-demo/config"
)

var DataBase *gorm.DB

func InitMysqlDB() {
	config := config.Config
	mySqlconfig := config.MysqlConfig
	host := mySqlconfig.Host
	port := mySqlconfig.Port
	username := mySqlconfig.Username
	password := mySqlconfig.PassWord
	database := mySqlconfig.Database
	args := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		username,
		password,
		host,
		port,
		database)
	//匹配gorm.Open函数
	var err error = nil
	DataBase, err = gorm.Open("mysql", args)
	if err != nil {
		panic("failed to connect database ,err:" + err.Error())
	}
	//log.Infof("connect database success,user:%s,database:%s", username, database)
}
func GetDB() *gorm.DB {
	return DataBase
}
func CloseDataBase() {
	DataBase.Close()
}
