package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/rsdel2007/proj/internal/model"
	"gorm.io/gorm"
)

func RegisterHandlers(db *gorm.DB, r *mux.Router) {
	r.HandleFunc("/videos", getVideosHandler(db)).Methods("GET")
	r.HandleFunc("/search", searchVideosHandler(db)).Methods("GET")
}

func getVideosHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}

		pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
		if err != nil || pageSize < 1 {
			http.Error(w, "Invalid pageSize parameter", http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pageSize

		videos := model.GetPaginatedVideos(db, offset, pageSize)
		if len(videos) == 0 {
			http.NotFound(w, r)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(videos)
	}
}

func searchVideosHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		query := r.URL.Query().Get("query")
		if query == "" {
			http.Error(w, "Invalid Query parameter", http.StatusBadRequest)
			return
		}
		page, err := strconv.Atoi(r.URL.Query().Get("page"))
		if err != nil || page < 1 {
			http.Error(w, "Invalid page parameter", http.StatusBadRequest)
			return
		}

		pageSize, err := strconv.Atoi(r.URL.Query().Get("pageSize"))
		if err != nil || pageSize < 1 {
			http.Error(w, "Invalid pageSize parameter", http.StatusBadRequest)
			return
		}

		offset := (page - 1) * pageSize

		videos, err := model.SearchVideos(db, query, offset, pageSize)
		if err != nil {
			http.Error(w, "Didn't get results for this query", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(videos)
	}

}
