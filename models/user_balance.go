package models

import (
	"time"
)

type UserBalance struct {
	ID 					int  `gorm:"primary_key" json:"id"`
	UserID 				int  `gorm:"column:user_id" json:"user_id"`
	Balance 			int	 `gorm:"column:balance"`
	BalanceAchieve  	int	 `gorm:"column:balance_achieve"`
	UserBalanceHistory 	[]UserBalanceHistory
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
}



