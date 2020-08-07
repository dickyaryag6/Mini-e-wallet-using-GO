package controllers

import (
	"ewallet/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"path"
)

func serveTemplate(data map[string]interface{}, w http.ResponseWriter, filename string) {
	filepath := path.Join("views", filename)
	tmpl, err := template.ParseFiles(filepath)
	// if failed to serve template
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.Execute(w, data)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func Register(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != "POST" { // jika method POST
			serveTemplate(nil, w, "register.html") //tampilkan page register.html
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
			hashedPassword, err := bcrypt.GenerateFromPassword([]byte(Password), bcrypt.DefaultCost)
			if err != nil {
				log.Fatal(err)
			}

			// insert user baru
			err = con.Create(&models.User{Email: Email, Username: Username, Password: string(hashedPassword)}).Error
			if err != nil {
				log.Fatal(err)
				return
			}

			fmt.Println(w, "Registrasi berhasil")
		} else {
			//jika username atau email sudah ada di database
			fmt.Println(w, "Username atau email sudah terdaftar")
		}
	}
}
