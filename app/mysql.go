package app

import (
	"go-pos/helper"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var err error

func ConnectMysql(connectionString string) {
	Instance, err = gorm.Open(mysql.Open(connectionString), &gorm.Config{})
	helper.PanicIfError(err)

	log.Println("Connected to Database Mysql...")
}
