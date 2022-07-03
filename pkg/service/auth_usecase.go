package service

import (
	"github.com/SimilarEgs/CRUD-TODO-LIST/internal/entity"
	"github.com/SimilarEgs/CRUD-TODO-LIST/utils"
)

// implementing authorization interface
type AuthService struct {
	repo RepositoryAuthorization
}

func NewAuthService(repo RepositoryAuthorization) *AuthService {
	return &AuthService{repo: repo}
}

// implementation of createUser method
func (s *AuthService) CreateUser(user entity.User) (int, error) {

	var err error

	user.Password, err = utils.GenHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	// pass user struct on repo layer
	return s.repo.CreateUser(user)
}
