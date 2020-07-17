package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

var (
	DB *gorm.DB
)

func MysqlInit() {
	sql,err :=gorm.Open("mysql","root:@tcp(127.0.0.1:3306)/ongorm?charset=utf8&parseTime=true")
	if err!=nil {
		fmt.Println(err.Error())
		return
	}
	DB=sql
	_ = check_table(&User{})
	_ = check_table(&Ticket{})
	_ = check_table(&Order{})
}

func check_table(typ interface{}) (err error) {
	if !DB.HasTable(typ) {
		dd:=DB.CreateTable(typ)
		err = dd.Error
	}
	if err != nil {
		log.Println(err)
	}
	return
}



