package routes

import (
	"ewallet/controllers"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	"net/http"
)

// GetRoutes : fungsi untuk mendapatkan semua routes yang ada
func GetRoutes(con *gorm.DB) *mux.Router {
	routes := mux.NewRouter().StrictSlash(false)

	// redirect ke route index
	routes.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/index", 302)
	})

	//tambahkan route di bawah ini
	routes.HandleFunc("/index", controllers.Index(con)) // jika route /index dipanggil, maka akan menjalankan Handler Index

	//autentikasi
	routes.HandleFunc("/register", controllers.Register(con)).Methods("POST")
	routes.HandleFunc("/login", controllers.Login(con)).Methods("POST")
	routes.HandleFunc("/logout", controllers.Logout(con))

	//transaksi wallet
	routes.HandleFunc("/wallet/", controllers.GetAllWallet(con)).Methods("GET")
	routes.HandleFunc("/wallet/{walletID}", controllers.GetWallet(con)).Methods("GET")
	routes.HandleFunc("/wallet/create", controllers.CreateNewWallet(con)).Methods("POST")
	routes.HandleFunc("/wallet/delete/{walletID}", controllers.DeleteWallet(con)).Methods("DELETE")
	routes.HandleFunc("/wallet/addbalance/{walletID}", controllers.AddBalance(con)).Methods("PUT")
	routes.HandleFunc("/wallet/substractbalance/{walletID}", controllers.SubstractBalance(con)).Methods("PUT")

	//transfer antar wallet


	return routes
}