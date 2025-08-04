package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"

	"sudoku/internal/models"
)

type PuzzleHandler struct {
	db *gorm.DB
}

func NewPuzzleHandler(db *gorm.DB) *PuzzleHandler {
	return &PuzzleHandler{db: db}
}

func (h *PuzzleHandler) GetPuzzles(w http.ResponseWriter, r *http.Request) {
	difficulty := r.URL.Query().Get("difficulty")
	limitStr := r.URL.Query().Get("limit")

	limit := 10 // default limit
	if limitStr != "" {
		if l, err := strconv.Atoi(limitStr); err == nil && l > 0 {
			limit = l
		}
	}

	query := h.db.Model(&models.Puzzle{})

	if difficulty != "" {
		// Validate difficulty
		switch models.Difficulty(difficulty) {
		case models.Easy, models.Medium, models.Hard:
			query = query.Where("difficulty = ?", difficulty)
		default:
			http.Error(w, "Invalid difficulty level", http.StatusBadRequest)
			return
		}
	}

	var puzzles []models.Puzzle
	if err := query.Limit(limit).Find(&puzzles).Error; err != nil {
		http.Error(w, "Failed to fetch puzzles", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(puzzles)
}
