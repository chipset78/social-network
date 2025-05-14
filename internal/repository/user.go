package repository

import (
	"context"
	"database/sql"
	"log"
	"social-network/internal/config"
	"social-network/internal/models"
	"social-network/internal/utils"
	"time"

	"github.com/google/uuid"
)

type UserRepository struct {
	db  *sql.DB
	cfg *config.Config
}

func NewUserRepository(db *sql.DB, cfg *config.Config) *UserRepository {
	return &UserRepository{
		db:  db,
		cfg: cfg,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *models.User) error {
	user.ID = uuid.New()
	user.CreatedAt = time.Now()

	query := `INSERT INTO users (id, first_name, second_name, birthdate, biography, city, password, created_at)
	          VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`

	_, err := r.db.ExecContext(ctx, query,
		user.ID,
		user.FirstName,
		user.LastName,
		user.BirthDate,
		user.Biography,
		user.City,
		user.Password,
		user.CreatedAt,
	)

	return err
}

func (r *UserRepository) GetByID(ctx context.Context, id uuid.UUID) (*models.User, error) {
	query := `SELECT id, first_name, second_name, birthdate, biography, city 
	          FROM users WHERE id = $1`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.BirthDate,
		&user.Biography,
		&user.City,
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) CheckCredentials(ctx context.Context, id uuid.UUID, password string) error {
	query := `SELECT id, password
	          FROM users WHERE id = $1`

	var user models.User
	err := r.db.QueryRowContext(ctx, query, id).Scan(
		&user.ID,
		&user.Password,
	)
	if err != nil {
		log.Printf("DB error: %v", err)
		return err
	}

	// Проверяем пароль с использованием соли из конфига
	if !utils.CheckPasswordHash(password, user.Password, r.cfg.PasswordSalt) {
		return sql.ErrNoRows
	}

	return nil
}
