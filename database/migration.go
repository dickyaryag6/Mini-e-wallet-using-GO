package database

import (
	"ewallet/models"
	"github.com/jinzhu/gorm"
)

func Migrate(con *gorm.DB) {
	con.DropTableIfExists(models.UserBalanceHistory{}, models.UserBalance{}, models.User{})
	con.AutoMigrate(models.User{}, models.UserBalance{}, models.UserBalanceHistory{})

	//add foreign key
	con.Model(&models.UserBalance{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	con.Model(&models.UserBalanceHistory{}).AddForeignKey("user_balance_id", "user_balances(id)", "CASCADE", "RESTRICT")
}