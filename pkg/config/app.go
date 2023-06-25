package config

import (
	"fmt"

	"github.com/jinzhu/gorm" // todo: switch to github.com/go-gorm/gorm?
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	fmt.Println("connecting to DB")
	d, err := gorm.Open("mysql", "root:myPassword@tcp(127.0.0.1:3306)/bookStore?charset=utf8&parseTime=True&loc=Local")
	// gorm.Open("mysql","userName:password/DBName?some-required-stuff-by-mysql")
	if err != nil {
		panic(err)
	}
	db = d
	fmt.Println("DB connected")
}

func GetDB() *gorm.DB {
	return db
}
