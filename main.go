package main

import (
	"ewallet/database"
	"ewallet/routes"
	"fmt"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/spf13/viper"
	"log"
	"net/http"
)

func GetEnvironmentVariable(key string) string {
	viper.SetConfigFile(".env")

	err := viper.ReadInConfig()

	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}

	value, ok := viper.Get(key).(string)

	if !ok {
		log.Fatalf("Invalid type assertion")
	}

	return value
}

func main() {

	//Inisialisasi Koneksi Database
	connection := database.DBConnection(
			GetEnvironmentVariable("USER"),
			GetEnvironmentVariable("PASSWORD"),
			GetEnvironmentVariable("HOST"),
			GetEnvironmentVariable("DATABASE"),
	)

	//migrasi database
	database.Migrate(connection)

	//melakukan seed untuk data awal di database
	database.Seed(connection)

	routes := routes.GetRoutes(connection)
	http.Handle("/", routes)

	PORT := GetEnvironmentVariable("PORT")
	fmt.Println(fmt.Sprintf("server started at localhost:%s", PORT))
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", PORT), nil))

	defer connection.Close()
}
