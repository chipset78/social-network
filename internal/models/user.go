package models

import (
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID `json:"id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"second_name"`
	BirthDate time.Time `json:"birthdate"`
	Biography string    `json:"biography"`
	City      string    `json:"city"`
	Password  string    `json:"-"`
	CreatedAt time.Time `json:"-"`
}

type RegisterRequest struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"second_name"`
	BirthDate string `json:"birthdate"`
	Biography string `json:"biography"`
	City      string `json:"city"`
	Password  string `json:"password"`
}

type LoginRequest struct {
	ID       string `json:"id"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Token string `json:"token"`
}

type RegisterResponse struct {
	UserID string `json:"user_id"`
}
