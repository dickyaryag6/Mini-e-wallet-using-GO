package controllers

import (
	"encoding/json"
	"ewallet/models"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type UpdateBalance struct {
	newBalance int
}

func getCookieValue(con *gorm.DB, r *http.Request) string {
	c := &http.Cookie{}
	if storedCookie, _ := r.Cookie("User_ID"); storedCookie != nil {
	c = storedCookie
	}
	UserId := c.Value

	return UserId
}

func JsonResponse(data interface{}, w http.ResponseWriter) {

	w.Header().Set("Content-Type", "application/json")

	err := json.NewEncoder(w).Encode(data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func GetWallet(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "GET" {
			//mendapatkan walletID dari url parameter
			urlParams := mux.Vars(r)
			walletID := urlParams["walletID"]

			var wallet models.UserBalance

			if err := con.First(&wallet, walletID).Error; gorm.IsRecordNotFoundError(err) {
				// record not found
				http.Error(w, "Wallet tidak ditemukan", http.StatusBadRequest)
				return
			}

			data := struct {wallet models.UserBalance} {wallet}

			JsonResponse(data, w)
		}
	}
}

func CreateNewWallet(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "POST" {
			//mendapatkan value dari cookie
			userId, _ := strconv.Atoi(getCookieValue(con, r))

			fmt.Println("new",userId)

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
			}

			JsonResponse(nil,  w)
		}
	}
}

func DeleteWallet(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "DELETE" {
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

			JsonResponse(nil, w)
		}
	}
}

func AddBalance(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
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
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if wallet.UserID != userId {
				http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			}

			//mendapatkan newBalance dari url parameters
			newBalance, _ := strconv.Atoi(strings.Split(r.URL.RawQuery, "=")[1])

			//mendapatkan nilai balance sebelumnya
			con.First(&wallet, walletID)
			oldBalance := wallet.Balance

			con.Model(&wallet).Where("id = ?", walletID).Update("balance", newBalance+oldBalance)

			data := struct {wallet models.UserBalance} {wallet}

			JsonResponse(data, w)
		}
	}
}

func SubstractBalance(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		if r.Method == "PUT" {
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
			userId, _ := strconv.Atoi(getCookieValue(con, r))
			if wallet.UserID != userId {
				http.Error(w, "Anda tidak punya akses", http.StatusBadRequest)
			}

			//mendapatkan newBalance dari url parameters
			newBalance, _ := strconv.Atoi(strings.Split(r.URL.RawQuery, "=")[1])

			//mendapatkan nilai balance sebelumnya
			con.First(&wallet, walletID)
			oldBalance := wallet.Balance

			if newBalance < oldBalance {
				con.Model(&wallet).Where("id = ?", walletID).Update("balance", oldBalance - newBalance)
				data := struct {wallet models.UserBalance} {wallet}
				JsonResponse(data, w)
			} else {
				JsonResponse(nil,w)
			}
		}
	}
}

