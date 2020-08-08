package models

import "time"

type BankBalance struct {
	ID 					int    `gorm:"primary_key" json:"id"`
	Balance 			int	   `gorm:"column:balance"`
	BalanceAchieve  	int	   `gorm:"column:balance_achieve"`
	BankCode			string `gorm:"column:code;unique"`
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
}

func (bb BankBalance) CheckBalance (newBalance int) bool {
	if bb.Balance >= newBalance {
		return true
	}
	return false
}
