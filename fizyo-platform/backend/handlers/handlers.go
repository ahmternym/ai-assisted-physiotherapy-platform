package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"fizyo-platform/models"
)

// GetPhysiotherapistsByCity - GET /api/physiotherapists?city={city}
// CORS header'ları artık main.go'daki corsMiddleware tarafından ekleniyor.
func GetPhysiotherapistsByCity(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	city := r.URL.Query().Get("city")
	if city == "" {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "city parametresi gerekli"})
		return
	}

	var result []models.Physiotherapist
	for _, p := range physiotherapistsData {
		if strings.EqualFold(p.City, city) {
			result = append(result, p)
			if len(result) == 5 {
				break
			}
		}
	}

	if result == nil {
		result = []models.Physiotherapist{}
	}

	json.NewEncoder(w).Encode(result)
}

// GetPhysiotherapistByID - GET /api/physiotherapists/{id}
func GetPhysiotherapistByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	// Trailing slash'ı temizle: /api/physiotherapists/5 veya /api/physiotherapists/5/
	path := strings.TrimSuffix(r.URL.Path, "/")
	parts := strings.Split(path, "/")
	idStr := parts[len(parts)-1]

	id, err := strconv.Atoi(idStr)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "geçersiz id"})
		return
	}

	for _, p := range physiotherapistsData {
		if p.ID == id {
			json.NewEncoder(w).Encode(p)
			return
		}
	}

	w.WriteHeader(http.StatusNotFound)
	json.NewEncoder(w).Encode(map[string]string{"error": "fizyoterapist bulunamadı"})
}
