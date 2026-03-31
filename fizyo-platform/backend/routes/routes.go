package routes

import (
	"net/http"
	"strings"

	"fizyo-platform/handlers"
)

// RegisterRoutes tüm API endpoint'lerini mux'a kaydeder.
// Sorun: Go'nun net/http mux'unda "/api/physiotherapists/" pattern'i
// hem /api/physiotherapists hem de /api/physiotherapists/1 gibi path'leri yakalar.
// Bu yüzden tek bir handler yazıp içinde path parse ederek ayırt ediyoruz.
func RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/api/physiotherapists", handlers.GetPhysiotherapistsByCity)
	mux.HandleFunc("/api/physiotherapists/", func(w http.ResponseWriter, r *http.Request) {
		// Trailing slash'ı temizle ve parçaları kontrol et
		path := strings.TrimSuffix(r.URL.Path, "/")
		parts := strings.Split(path, "/")

		// /api/physiotherapists/{id} → parts = ["", "api", "physiotherapists", "{id}"]
		if len(parts) == 4 && parts[3] != "" {
			handlers.GetPhysiotherapistByID(w, r)
			return
		}

		// Sadece /api/physiotherapists/ gelirse liste endpoint'ine yönlendir
		handlers.GetPhysiotherapistsByCity(w, r)
	})
}
