package repository

import (
	"fmt"

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

	var id int

	// query for creation user in db
	query := fmt.Sprintf("INSERT INTO %s (name, username, hashed_password) VALUES ($1, $2, $3) RETURNING id", usersTable)

	// execution of the sql statement
	row := r.db.QueryRow(query, user.Name, user.Username, user.Password)

	// storing user id, checks if any error
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}
