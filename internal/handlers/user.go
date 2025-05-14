package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"social-network/internal/config"
	"social-network/internal/models"
	"social-network/internal/repository"
	"social-network/internal/utils"
	"time"

	"github.com/google/uuid"
)

type UserHandler struct {
	repo *repository.UserRepository
	cfg  *config.Config
}

func NewUserHandler(repo *repository.UserRepository, cfg *config.Config) *UserHandler {
	return &UserHandler{
		repo: repo,
		cfg:  cfg,
	}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var req models.RegisterRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	hashedPassword, err := utils.HashPassword(req.Password, h.cfg.PasswordSalt)
	if err != nil {
		http.Error(w, "Failed to hash password", http.StatusInternalServerError)
		return
	}

	birthDate, err := time.Parse("2006-01-02", req.BirthDate)
	if err != nil {
		http.Error(w, "Invalid birthdate format", http.StatusBadRequest)
		return
	}

	user := &models.User{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		BirthDate: birthDate,
		Biography: req.Biography,
		City:      req.City,
		Password:  hashedPassword,
	}

	if err := h.repo.Create(r.Context(), user); err != nil {
		http.Error(w, "Failed to create user", http.StatusInternalServerError)
		return
	}

	response := models.RegisterResponse{UserID: user.ID.String()}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func (h *UserHandler) GetUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := uuid.Parse(idStr)
	if err != nil || id == uuid.Nil {
		http.Error(w, "Invalid user ID format", http.StatusBadRequest)
		return
	}

	user, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "User not found", http.StatusNotFound)
		} else {
			http.Error(w, "Database error", http.StatusInternalServerError)
		}
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(user)
}
