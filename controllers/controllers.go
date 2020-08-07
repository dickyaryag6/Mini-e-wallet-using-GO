package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

// Index : Handler utama
func Index(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//mendapatkan username user

		c := &http.Cookie{}
		if storedCookie, _ := r.Cookie("Username"); storedCookie != nil {
			c = storedCookie
		}

		fmt.Println("Hello",c.Value)

		//data := map[string]string{
		//	"username": c.,
		//	"message":  "Index Page !",
		//}

		//fmt.Println("Hello")
	}
}
