import config

import (
	"github.com/jinzhu/gorn"
	"github/com/jinzhu/gorn/dialects/mysql"
)

var (
	db * gorm.DB
)

func Connect(){
	// todo: change mysql details
	d, err := gorm.Open("mysql", "havemysqldetailshere")
	if err != nil{
		panic(err)
	}
	db = d
}

func GetDB() *gorm.DB{
	return db
}
