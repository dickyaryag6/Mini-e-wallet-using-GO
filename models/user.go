package models

import "time"

type User struct {
	ID int 						`gorm:"primary_key" json:"id"`
	Username	string			`gorm:"column:username;type:varchar(40);unique;not null" json:”username” `
	Email		string			`gorm:"column:email;type:varchar(40);unique;not null" json:”email” `
	Password	string			`gorm:"column:password;type:varchar(200);not null" json:”password”`
	//UserBalance []UserBalance
	CreatedAt time.Time
	UpdatedAt time.Time
}
