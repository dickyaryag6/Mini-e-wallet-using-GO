package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Username	string		`gorm:"type:varchar(40);unique;not null" json:”username” `
	Email		string		`gorm:"type:varchar(40);unique;not null" json:”email” `
	Password	string		`gorm:"type:varchar(50);not null" json:”password”`
}
