package database

import (
	"ewallet/models"
	"github.com/jinzhu/gorm"
)

func Migrate(con *gorm.DB) {
	con.DropTableIfExists(models.UserBalanceHistory{}, models.UserBalance{}, models.User{}, models.BankBalanceHistory{}, models.BankBalance{})
	con.AutoMigrate(models.User{}, models.UserBalance{}, models.UserBalanceHistory{}, models.BankBalance{}, models.BankBalanceHistory{})

	//add foreign key
	con.Model(&models.UserBalance{}).AddForeignKey("user_id", "users(id)", "CASCADE", "RESTRICT")
	con.Model(&models.UserBalanceHistory{}).AddForeignKey("user_balance_id", "user_balances(id)", "CASCADE", "RESTRICT")
	con.Model(&models.BankBalanceHistory{}).AddForeignKey("bank_balance_id", "bank_balances(id)", "CASCADE", "RESTRICT" )
}