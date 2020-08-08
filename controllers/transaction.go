package controllers

import (
	"encoding/json"
	"ewallet/models"
	"github.com/gorilla/mux"
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

				//data := struct {message string;wallet []models.UserBalance} {"Berhasil", wallet}
				//data := struct {message string} {"Berhasil"}
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
				urlParams := mux.Vars(r)
				walletID := urlParams["walletID"]

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
					JsonResponse(nil,  w)
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
				urlParams := mux.Vars(r)
				walletID := urlParams["walletID"]

				//cek jika wallet dengan id tersebut ada
				var wallet models.UserBalance

				if err := con.First(&wallet, walletID).Error; gorm.IsRecordNotFoundError(err) {
					// record not found
					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
					return
				}

				con.Delete(&wallet, walletID)

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
				urlParams := mux.Vars(r)
				walletID := urlParams["walletID"]

				var wallet models.UserBalance

				//cek jika wallet dengan id tersebut ada
				if err := con.First(&wallet, walletID).Error; gorm.IsRecordNotFoundError(err) {
					// record not found
					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
					return
				}

				//cek jika user_id dari wallet sama dengan id dari user yang login
				if wallet.UserID != userId {
					http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
				}

				//mendapatkan newBalance dari url parameters
				newBalance, _ := strconv.Atoi(r.URL.Query()["newbalance"][0])

				//mendapatkan nilai balance sebelumnya
				con.First(&wallet, walletID)
				oldBalance := wallet.Balance

				//update nilai balance dan balance_achieve
				con.Model(&wallet).Updates(map[string]interface{}{"balance": newBalance+oldBalance, "balance_achieve": newBalance+oldBalance})

				JsonResponse(wallet, w)
				return
			}

			//tidak punya akses
			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			return
		}
	}
}

func SubstractBalance(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
			//cek userid dari cookie
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if userId != 0 {
				//mendapatkan walletID dari url parameter
				urlParams := mux.Vars(r)
				walletID := urlParams["walletID"]

				var wallet models.UserBalance

				//cek jika wallet dengan id tersebut ada
				if err := con.First(&wallet, walletID).Error; gorm.IsRecordNotFoundError(err) {
					// record not found
					http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
					return
				}

				//cek jika user_id dari wallet sama dengan id dari user yang login
				if wallet.UserID != userId {
					http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
				}

				//mendapatkan newBalance dari url parameters
				newBalance, _ := strconv.Atoi(r.URL.Query()["newbalance"][0])

				//mendapatkan nilai balance sebelumnya
				con.First(&wallet, walletID)
				oldBalance := wallet.Balance

				if newBalance < oldBalance {
					con.Model(&wallet).Where("id = ?", walletID).Update("balance", oldBalance - newBalance)
					JsonResponse(wallet, w)
					return
				} else {
					http.Error(w, "Saldo tidak mencukupi", http.StatusBadRequest)
					return
				}
			}

			//tidak punya akses
			http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			return
		}
	}
}

