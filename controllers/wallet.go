package controllers

import (
	"encoding/json"
	"ewallet/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"time"
)

func getCookieValue(con *gorm.DB, r *http.Request) string {
	c := &http.Cookie{}
	if storedCookie, _ := r.Cookie("User_ID"); storedCookie != nil {
	c = storedCookie
	}
	UserId := c.Value

	return UserId
}

func JsonResponse(data interface{}, w http.ResponseWriter) {
	result, err := json.Marshal(data)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(result)
}


func GetAllWallet(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "GET" {
			//cek userid dari cookie
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if userId != 0 {
				//get semua wallet dengan user_id sama dengan userId
				wallet := make([]models.UserBalance, 0)

				if err := con.Where(&models.UserBalance{UserID: userId}).Find(&wallet).Error; gorm.IsRecordNotFoundError(err) {
					// record not found
					http.Error(w, "Anda belum mempunyai wallet", http.StatusBadRequest)
					return
				}

				JsonResponse(wallet, w)
				return
			}

			//tidak punya akses
			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			return
		}
	}
}

func GetWallet(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			//cek userid dari cookie
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if userId != 0 {
				//get semua wallet dengan user_id sama dengan userId
				//mendapatkan walletID dari url parameter
				walletID, _ := strconv.Atoi(r.URL.Query()["walletid"][0])

				var wallet models.UserBalance

				if err := con.First(&wallet, walletID).Error; gorm.IsRecordNotFoundError(err) {
					// record not found
					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
					return
				}

				//cek jika user punya akses terhadap wallet
				if wallet.UserID == userId {
					//punya akses
					JsonResponse(wallet, w)
					return
				}
			}
			//tidak punya akses
			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			return
		}
	}
}

func CreateNewWallet(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			//cek userid dari cookie
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if userId != 0 {
				var wallet = models.UserBalance{
					UserID:             userId,
					Balance:            0,
					BalanceAchieve:     0,
					CreatedAt:          time.Time{},
					UpdatedAt:          time.Time{},
				}

				err := con.Create(&wallet).Error
				if err != nil {
					http.Error(w, "Membuat Wallet baru gagal", http.StatusBadRequest)
					return
				}

				w.Write([]byte("Membuat Wallet baru berhasil"))
				return
			}
			//tidak punya akses
			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			return
		}
	}
}


func DeleteWallet(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "DELETE" {
			//cek userid dari cookie
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if userId != 0 {
				//mendapatkan walletID dari url parameter
				walletID, _ := strconv.Atoi(r.URL.Query()["walletid"][0])

				//cek jika wallet dengan id tersebut ada
				var wallet models.UserBalance

				if err := con.First(&wallet, walletID).Error; gorm.IsRecordNotFoundError(err) {
					// record not found
					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
					return
				}

				//cek jika user_id dari wallet sama dengan id dari user yang login
				if wallet.UserID != userId {
					http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
					return
				}

				err := con.Delete(&wallet, walletID).Error
				if err != nil {
					http.Error(w, "Menghapus Wallet baru gagal", http.StatusBadRequest)
					return
				}

				w.Write([]byte("Menghapus wallet berhasil"))
				return
			}
			//tidak punya akses
			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			return
		}
	}
}

func AddBalance(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			//cek userid dari cookie
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if userId != 0 {
				//mendapatkan walletID dari url parameter
				walletID, _ := strconv.Atoi(r.URL.Query()["walletid"][0])
				transactionType := r.URL.Query()["type"][0]
				bankCode := r.URL.Query()["code"][0]

				var wallet models.UserBalance

				//cek jika wallet dengan id tersebut ada
				if err := con.First(&wallet, walletID).Error; gorm.IsRecordNotFoundError(err) {
					// record not found//
					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
					return
				}

				//cek jika user_id dari wallet sama dengan id dari user yang login
				if wallet.UserID != userId {
					http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
					return
				}

				//mendapatkan newBalance dari url parameters
				newBalance, _ := strconv.Atoi(r.URL.Query()["newbalance"][0])

				//cek jika bank ada
				var bank models.BankBalance
				if err := con.Where("code = ?", bankCode).First(&bank).Error; err != nil {
					http.Error(w, "Bank tidak ditemukan", http.StatusBadRequest)
					return
				}

				//cek jika saldo bank cukup
				if !bank.CheckBalance(newBalance) {
					http.Error(w, "Silahkan pilih Bank lain", http.StatusBadRequest)
					return
				}

				//mendapatkan nilai balance sebelumnya
				con.First(&wallet, walletID)
				oldBalance := wallet.Balance

				//update nilai balance dan balance_achieve
				con.Model(&wallet).Updates(map[string]interface{}{"balance": newBalance+oldBalance, "balance_achieve": newBalance+oldBalance})

				CreateNewUserBalanceHistory(con, walletID, oldBalance, newBalance+oldBalance, "TopUp", transactionType, r)

				//kurangi balance Bank
				oldBankBalance := bank.Balance
				con.Model(&bank).Updates(map[string]interface{}{"balance": oldBankBalance-newBalance})

				CreateNewBankBalanceHistory(con, bank.ID, bank.Balance, newBalance,"TopUp",transactionType, r)

				JsonResponse(wallet, w)
				return
			}

			//tidak punya akses
			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			return
		}
	}
}

//func SubstractBalance(con *gorm.DB) http.HandlerFunc {
//	return func(w http.ResponseWriter, r *http.Request) {
//
//		if r.Method == "PUT" {
//			//cek userid dari cookie
//			userId, _ := strconv.Atoi(getCookieValue(con, r))
//			if userId != 0 {
//				//mendapatkan walletID dari url parameter
//				urlParams := mux.Vars(r)
//				walletID, _ := strconv.Atoi(r.URL.Query()["walletid"][0])
//				transactionType := r.URL.Query()["type"][0]
//				bankCode := urlParams["code"]
//
//				var wallet models.UserBalance
//
//				//cek jika wallet dengan id tersebut ada
//				if err := con.First(&wallet, walletID).Error; gorm.IsRecordNotFoundError(err) {
//					// record not found
//					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
//					return
//				}
//
//				//cek jika user_id dari wallet sama dengan id dari user yang login
//				if wallet.UserID != userId {
//					http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
//				}
//
//				//cek jika bank ada
//				var bank models.BankBalance
//				if err := con.Where("code = ?", bankCode).First(&bank).Error; err != nil {
//					http.Error(w, "Bank tidak ditemukan", http.StatusBadRequest)
//					return
//				}
//
//				//mendapatkan newBalance dari url parameters
//				newBalance, _ := strconv.Atoi(r.URL.Query()["newbalance"][0])
//
//				//mendapatkan nilai balance sebelumnya
//				con.First(&wallet, walletID)
//				oldBalance := wallet.Balance
//
//				if newBalance < oldBalance {
//					con.Model(&wallet).Updates(map[string]interface{}{"balance": oldBalance-newBalance, "balance_achieve": oldBalance-newBalance})
//
//					CreateNewUserBalanceHistory(con, userId, oldBalance, oldBalance-newBalance, "Transfer", transactionType, r)
//
//					JsonResponse(wallet, w)
//					return
//				} else {
//					http.Error(w, "Saldo tidak mencukupi", http.StatusBadRequest)
//					return
//				}
//			}
//
//			//tidak punya akses
//			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
//			return
//		}
//	}
//}


func Transfer (con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			//cek userid dari cookie
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if userId != 0 {

				fromWalletID, _ := strconv.Atoi(r.URL.Query()["fromwallet"][0])
				toWalletID, _ := strconv.Atoi(r.URL.Query()["towallet"][0])
				transactionType := r.URL.Query()["type"][0]

				transferBalance, _ := strconv.Atoi(r.URL.Query()["balance"][0])

				//get fromWallet
				var fromWallet models.UserBalance
				if err := con.First(&fromWallet, fromWalletID).Error; gorm.IsRecordNotFoundError(err) {
					// record not found
					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
					return
				}

				//cek balance fromWallet


				//get toWallet
				var toWallet models.UserBalance
				if err := con.First(&toWallet, toWalletID).Error; gorm.IsRecordNotFoundError(err) {
					// record not found
					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
					return
				}

				fmt.Println(toWallet.Balance, fromWallet.Balance)

				//transfer

				oldFromBalance := fromWallet.Balance
				if transferBalance > fromWallet.Balance {
					http.Error(w, "Saldo tidak cukup", http.StatusBadRequest)
					return
				}
				//kurangi balance fromWallet
				oldToBalance := toWallet.Balance

				con.Model(&fromWallet).Updates(map[string]interface{}{"balance": oldFromBalance-transferBalance})

				//tambahkan balance toWallet
				con.Model(&toWallet).Updates(map[string]interface{}{"balance": transferBalance+oldToBalance, "balance_achieve": transferBalance+oldToBalance})

				//tambahkan history
				CreateNewUserBalanceHistory(con, fromWalletID, oldFromBalance, oldFromBalance-transferBalance, "Transfer", transactionType, r)
				CreateNewUserBalanceHistory(con, toWalletID, oldToBalance, oldToBalance-transferBalance, "Transfer", transactionType, r)

				w.Write([]byte("Transfer Berhasil"))
				return
			}
			//tidak punya akses
			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			return
		}
	}
}

func CreateNewUserBalanceHistory(con *gorm.DB, userId, balanceBefore, balanceAfter int, activity, transactionType string, r *http.Request) error {
	//get IP
	IPAddr := ""
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		IPAddr = forwarded
	}
	IPAddr = r.RemoteAddr

	var balanceHistory = models.UserBalanceHistory{
		UserBalanceID:   userId,
		BalanceBefore:   balanceBefore,
		BalanceAfter:    balanceAfter,
		Activity:        activity,
		TransactionType: transactionType,
		IP:              IPAddr,
		Location:        "",
		UserAgent:       r.UserAgent(),
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

