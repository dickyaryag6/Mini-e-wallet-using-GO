package database

import (
	"ewallet/models"
	"github.com/jinzhu/gorm"
	"ewallet/controllers"
	"time"
)

func Seed(con *gorm.DB) {
	//seed data user
	hash1, _ := controllers.HashPassword("123456") //user 1
	hash2, _ := controllers.HashPassword("23456")  //user 2
	hash3, _ := controllers.HashPassword("privy")  //user 3

	con.Create(&models.User{Email: "abc@gmail.com", Username: "abc", Password: hash1})
	con.Create(&models.User{Email: "dicky@gmail.com", Username: "dicky", Password: hash2})
	con.Create(&models.User{Email: "123@gmail.com", Username: "123", Password: hash3})

	//seed data wallet
	wallet1 := models.UserBalance{UserID:1, Balance: 0, BalanceAchieve:0, CreatedAt:time.Time{}, UpdatedAt:time.Time{}}
	wallet2 := models.UserBalance{UserID:2, Balance: 0, BalanceAchieve:0, CreatedAt:time.Time{}, UpdatedAt:time.Time{}}
	wallet3 := models.UserBalance{UserID:3, Balance: 0, BalanceAchieve:0, CreatedAt:time.Time{}, UpdatedAt:time.Time{}}

	con.Create(&wallet1);con.Create(&wallet2);con.Create(&wallet3)

	//seed data bank
	bank1 := models.BankBalance{Balance: 500000, BalanceAchieve: 500000, BankCode: "002", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}
	bank2 := models.BankBalance{Balance: 700000, BalanceAchieve: 700000, BankCode: "008", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}
	bank3 := models.BankBalance{Balance: 900000, BalanceAchieve: 900000, BankCode: "009", CreatedAt: time.Time{}, UpdatedAt: time.Time{}}

	con.Create(&bank1);con.Create(&bank2);con.Create(&bank3);

}
