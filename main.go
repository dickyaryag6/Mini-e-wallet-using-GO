package main

import (
	"ewallet/routes"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
	"log"
	"net/http"
)

func DBConnection() (*gorm.DB) {

	db, _ := gorm.Open("sqlite3", "/tmp/gorm.db")
	defer db.Close()

	return db
}

func main() {

	//Inisialisasi Koneksi Database
	connection := DBConnection()

	routes := routes.GetRoutes(connection)
	http.Handle("/", routes)

	fmt.Println("server started at localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
