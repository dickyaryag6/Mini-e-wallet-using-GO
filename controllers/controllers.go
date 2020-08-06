package controllers

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"net/http"
)

// Index : Handler utama
func Index(con *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintln(w, "Hello World")
	}
}
