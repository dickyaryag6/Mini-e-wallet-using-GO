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
		if storedCookie, _ := r.Cookie("User_ID"); storedCookie != nil {
			c = storedCookie
		}

		fmt.Println(c.Value)

	}
}
