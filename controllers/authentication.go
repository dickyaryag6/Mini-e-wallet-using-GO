package controllers

import (
	_ "crypto/sha1"
	"ewallet/models"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}


func Register(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" { // jika method POST
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		//mendapatkan data dari form
		Username := r.FormValue("username")
		Email := r.FormValue("email")
		Password := r.FormValue("password")


		var user models.User

		//cek apakah username sudah ada
		err1 := con.First(&user, "username = ?", Username).Error

		//cek apakah email sudah ada
		err2 := con.First(&user, "email = ?", Email).Error

		//jika username dan email tidak ada di database
		if err1 != nil && err2 != nil {

			//hash password
			hash, _ := HashPassword(Password)

			// insert user baru
			err := con.Create(&models.User{Email: Email, Username: Username, Password: hash}).Error
			if err != nil {
				log.Fatal(err)
				return
			}

			http.Error(w, "Registrasi Berhasil", http.StatusOK)
		} else {
			//jika username atau email sudah ada di database
			http.Error(w, "Username atau Email sudah terdaftar", http.StatusBadRequest)
		}
	}
}


func Login(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//cek username di session
		c := &http.Cookie{}
		if storedCookie, _ := r.Cookie("Username"); storedCookie != nil {
			c = storedCookie
		}
		if c.Value != "" {
			http.Error(w, "Anda belum login", http.StatusBadRequest)
			return
		}


		if r.Method != "POST" { // jika method POST
			http.Error(w, "", http.StatusBadRequest)
			return
		}

		//mendapatkan data dari form
		Username := r.FormValue("username")
		Password := r.FormValue("password")

		var user models.User

		// query user dari database
		con.First(&user, "username = ?", Username)

		//compare password
		match := CheckPasswordHash(Password, user.Password)

		if match {
			// login berhasil
			c = &http.Cookie{}
			c.Name = "User_ID"
			c.Value = strconv.Itoa(user.ID)
			c.Expires = time.Now().Add(5 * time.Minute)
			http.SetCookie(w, c)
			http.Error(w, "Login Berhasil", http.StatusOK)
		} else {
			// login gagal
			http.Error(w, "Username atau Password salah", http.StatusBadRequest)

		}
	}
}

func Logout(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		c := &http.Cookie{}
		c.Name = "User_ID"
		c.Expires = time.Unix(0, 0)
		c.MaxAge = -1
		http.SetCookie(w, c)

		http.Error(w, "Logout Berhasil", http.StatusOK)
	}
}
