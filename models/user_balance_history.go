package models

import "time"

type UserBalanceHistory struct {
	ID int 						`gorm:"primary_key" json:"id"`
	UserBalanceID	int  		`gorm:"column:user_balance_id"` //`gorm:"foreignkey:ID"`
	BalanceBefore	int			`gorm:"column:balance_before;not null;default:0;type:int"`
	BalanceAfter    int			`gorm:"column:balance_after;not null;default:0;type:int"`
	Activity		string		`gorm:"column:activity;type:varchar(50)"`
	TransactionType	string  	`gorm:"column:transaction_type;type:varchar(50)"` //enum(credit, debit)
	IP				string		`gorm:"column:ip;type:varchar(50)"`
	Location		string		`gorm:"column:location;type:varchar(50)"`
	UserAgent		string		`gorm:"column:user_agent;type:varchar(50)"`
	Author			string		`gorm:"column:author;type:varchar(50)"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (ubh UserBalanceHistory) CreateNewTransactionHistory(walletID int) {

}
