package controllers

import (
	_ "crypto/sha1"
	"ewallet/models"
	"fmt"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"html/template"
	"log"
	"net/http"
	"path"
	"time"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	fmt.Print("hash awal :", string(bytes), "\n")
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

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
			hash, _ := HashPassword(Password)

			// insert user baru
			err := con.Create(&models.User{Email: Email, Username: Username, Password: hash}).Error
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


func Login(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		//cek username di session
		c := &http.Cookie{}
		if storedCookie, _ := r.Cookie("Username"); storedCookie != nil {
			c = storedCookie
		}
		if c.Value != "" {
			http.Redirect(w, r, "/", 302)
		}


		if r.Method != "POST" { // jika method POST
			http.ServeFile(w, r, "login.html") //tampilkan page login.html
			return
		}

		//mendapatkan data dari form
		Username := r.FormValue("username")
		Password := r.FormValue("password")

		var user models.User

		// query user dari database
		con.First(&user, "username = ?", Username)

		if user.Username == "" {
			fmt.Println(w, "Username atau Password salah")
			http.Redirect(w, r, "/", 302)
			return
		}

		//compare password
		match := CheckPasswordHash(Password, user.Password)

		if match {
			// login berhasil
			c = &http.Cookie{}
			c.Name = "Username"
			c.Value = Username
			c.Expires = time.Now().Add(5 * time.Minute)
			http.SetCookie(w, c)
			http.Redirect(w, r, "/", 302)
		} else {
			// login gagal
			fmt.Println(w, "Username atau Password salah")
			http.Redirect(w, r, "/login", 302)
		}
	}
}
