package repository

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/jmoiron/sqlx"
)

// implementing reposiory interface
type AuthRepository struct {
	db *sqlx.DB
}

func NewAuthRepository(db *sqlx.DB) *AuthRepository {
	return &AuthRepository{db: db}
}

func (r *AuthRepository) CreateUser(user entity.User) (int, error) {
	return 0, nil
}
