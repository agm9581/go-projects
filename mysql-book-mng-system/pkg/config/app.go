package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

func Connect() {
	d, err := gorm.Open("mysql", "root:password@/book_store?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		fmt.Println("Found an error connecting to the database")
		panic(err)
	}

	db = d

}

func GetDB() *gorm.DB {
	return db
}
