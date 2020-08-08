package controllers

import (
	"ewallet/models"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)


func GetAllBank (con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			banks := make([]models.BankBalance, 0)
			if err := con.Find(&banks).Error; gorm.IsRecordNotFoundError(err) {
				// record not found
				http.Error(w, "Anda belum mempunyai wallet", http.StatusBadRequest)
				return
			}

			JsonResponse(banks, w)
			return
		}
	}
}

func CreateNewBank (con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {

			bankCode := r.URL.Query()["code"][0]
			initialBalance, _ := strconv.Atoi(r.URL.Query()["balance"][0])

			//cek jika bank code sudah ada
			var bank1 models.BankBalance
			if err := con.Where("code = ?", bankCode).First(&bank1).Error; gorm.IsRecordNotFoundError(err) {
				// record not found
				http.Error(w, "Bank sudah terdaftar", http.StatusBadRequest)
				return
			}

			var bank2 = models.BankBalance{
				Balance:        initialBalance,
				BalanceAchieve: initialBalance,
				BankCode:       bankCode,
				CreatedAt:      time.Time{},
				UpdatedAt:      time.Time{},
			}

			err := con.Create(&bank2).Error
			if err != nil {
				http.Error(w, "Membuat Bank baru gagal", http.StatusBadRequest)
				return
			}
			w.Write([]byte("Membuat Bank baru berhasil"))
		}
	}
}

func AddBankBalance (con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			//bankcode
			bankCode := r.URL.Query()["bankcode"][0]
			//newbalance
			newBalance, _ := strconv.Atoi(r.URL.Query()["newbalance"][0])

			var bank models.BankBalance

			//cek jika bank dengan code tersebut ada
			if err := con.Where("code = ?", bankCode).First(&bank).Error; err != nil {
				http.Error(w, "Bank tidak ditemukan", http.StatusBadRequest)
				return
			}

			oldBalance := bank.Balance

			con.Model(&bank).Updates(map[string]interface{}{"balance": newBalance+oldBalance, "balance_achieve": newBalance+oldBalance})

			w.Write([]byte("Penambahan saldo bank berhasil"))
		}
	}
}

func SubstractBankBalance (con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			//bankcode
			bankCode := r.URL.Query()["bankcode"][0]
			//newbalance
			newBalance, _ := strconv.Atoi(r.URL.Query()["newbalance"][0])

			var bank models.BankBalance

			//cek jika bank dengan code tersebut ada
			if err := con.Where("code = ?", bankCode).First(&bank).Error; err != nil {
				http.Error(w, "Bank tidak ditemukan", http.StatusBadRequest)
				return
			}

			oldBalance := bank.Balance

			con.Model(&bank).Updates(map[string]interface{}{"balance": newBalance-oldBalance, "balance_achieve": newBalance-oldBalance})

			w.Write([]byte("Pengurangan saldo bank berhasil"))
		}
	}
}

func CreateNewBankBalanceHistory(con *gorm.DB, bankID, oldBalance, newBalance int, activity, transactionType string, r *http.Request) error {

	//get IP
	IPAddr := ""
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		IPAddr = forwarded
	}
	IPAddr = r.RemoteAddr

	var balanceHistory = models.BankBalanceHistory{
		BankBalanceID:   bankID,
		BalanceBefore:   oldBalance,
		BalanceAfter:    oldBalance+newBalance,
		Activity:        activity,
		TransactionType: transactionType,
		IP:              IPAddr,
		Location:        "",
		UserAgent:        r.UserAgent(),
		Author:          "",
		CreatedAt:       time.Time{},
		UpdatedAt:       time.Time{},
	}

	err := con.Create(&balanceHistory).Error
	if err != nil {
		return err
	}

	return nil
}