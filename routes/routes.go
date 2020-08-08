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

	//wallet
	routes.HandleFunc("/wallet/all", controllers.GetAllWallet(con)).Methods("GET")
	routes.HandleFunc("/wallet", controllers.GetWallet(con)).Methods("GET")
	routes.HandleFunc("/wallet/create", controllers.CreateNewWallet(con)).Methods("POST")
	routes.HandleFunc("/wallet/delete", controllers.DeleteWallet(con)).Methods("DELETE")
	routes.HandleFunc("/wallet/addbalance", controllers.AddBalance(con)).Methods("PUT")
	//routes.HandleFunc("/wallet/substractbalance", controllers.SubstractBalance(con)).Methods("PUT")

	//transfer antar wallet
	routes.HandleFunc("/wallet/transfer", controllers.Transfer(con)).Methods("POST")


	//bank
	routes.HandleFunc("/bank/all", controllers.GetAllBank(con)).Methods("GET")
	routes.HandleFunc("/bank/create", controllers.CreateNewBank(con)).Methods("POST")
	routes.HandleFunc("/bank/addbalance", controllers.AddBankBalance(con)).Methods("PUT")
	routes.HandleFunc("/bank/substractbalance", controllers.SubstractBankBalance(con)).Methods("PUT")




	return routes
}