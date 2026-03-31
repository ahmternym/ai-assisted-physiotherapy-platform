package main

import (
	"fmt"
	"log"
	"net/http"

	"fizyo-platform/routes"
)

// corsMiddleware tüm isteklere CORS başlıkları ekler ve OPTIONS preflight'larını yanıtlar.
func corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		// Preflight isteği ise 204 döndür, devam etme
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func main() {
	mux := http.NewServeMux()
	routes.RegisterRoutes(mux)

	// CORS middleware'i tüm mux'a sar
	handler := corsMiddleware(mux)

	port := ":8080"
	fmt.Println("🚀 Fizyo Platform API başlatıldı →", port)
	log.Fatal(http.ListenAndServe(port, handler))
}
