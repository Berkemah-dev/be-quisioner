package main

import (
	"log"
	"net/http"

	"github.com/Berkemah-dev/be-quisioner/config"
	"github.com/Berkemah-dev/be-quisioner/routes"
)

func main() {
	config.InitConfig() // Memuat konfigurasi dari .env
	router := routes.SetupRoutes()

	// Middleware untuk mengatur CORS
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if config.SetAccessControlHeaders(w, r) {
			return // Jika ini adalah permintaan preflight, hentikan eksekusi
		}
		router.ServeHTTP(w, r) // Lanjutkan ke rute yang sesuai
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
