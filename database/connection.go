package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
)

func DBConnection(user, password, host, database string) *gorm.DB {

	dbUrl := fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", user, password, host, database)

	db, err := gorm.Open("mysql", dbUrl)
	if err != nil {
		log.Fatalln(err)
	}

	return db
}
