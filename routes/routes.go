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


	return routes
}