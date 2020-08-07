package database

import (
	"ewallet/models"
	"github.com/jinzhu/gorm"
)

func Migrate(con *gorm.DB) {
	con.AutoMigrate(models.User{})
}