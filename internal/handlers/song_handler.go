package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"testApp/internal/models"
	"testApp/internal/services"
)

type SongHandler struct {
    Service *services.SongService
}

// добавить документацию swagger
func (h *SongHandler) GetSongs(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
        return
	}
    page, _ := strconv.Atoi(r.URL.Query().Get("page"))
    if page == 0 {
        page = 1
    }
    limit, _ := strconv.Atoi(r.URL.Query().Get("limit"))
    if limit == 0 {
        limit = 10
    }
    filters := map[string]string{
        "group": r.URL.Query().Get("group"),
        "song":  r.URL.Query().Get("song"),
    }

    songs, err := h.Service.GetSongs(page, limit, filters)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error getting songs: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(songs)
}

// добавить swagger документацию
func (h *SongHandler) AddSong(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid method", http.StatusMethodNotAllowed)
		return
	}
    var song models.Song
    if err := json.NewDecoder(r.Body).Decode(&song); err != nil {
        http.Error(w, fmt.Sprintf("Error decoding request body: %v", err), http.StatusBadRequest)
        return
    }

    addedSong, err := h.Service.AddSong(song)
    if err != nil {
        http.Error(w, fmt.Sprintf("Error adding song: %v", err), http.StatusInternalServerError)
        return
    }

    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(addedSong)
}

// Другие обработчики для PATCH, DELETE и получения текста
