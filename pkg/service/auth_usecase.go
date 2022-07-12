package service

import (
	"github.com/SimilarEgs/go-todo-app/internal/entity"
	"github.com/SimilarEgs/go-todo-app/utils"
)

// implementing authorization interface
type AuthService struct {
	repo RepositoryAuthorization
}

func NewAuthService(repo RepositoryAuthorization) *AuthService {
	return &AuthService{repo: repo}
}

// implementation of CreateUser method
func (s *AuthService) CreateUser(user entity.User) (int64, error) {

	var err error

	user.Password, err = utils.GenHashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	// pass user struct on repo layer
	return s.repo.CreateUser(user)
}

// implementation of GenerateToken method
func (s *AuthService) GenerateToken(username, password string) (string, error) {

	// get user from db
	user, err := s.repo.GetUser(username)
	if err != nil {
		return "", err
	}

	// varify password
	err = utils.CheckPasswordMatch(password, user.Password)
	if err != nil {
		return "", err
	}

	// genereate token
	token, err := utils.CreateToken(username, user.Id)
	if err != nil {
		return "", err
	}

	return token, err
}

// implementation of ParseToken method
func (s *AuthService) ParseToken(token string) (int64, error) {

	// parsing access token
	clamis, err := utils.VerifyJWT(token)
	if err != nil {
		return 0, err
	}

	return clamis.ID, err
}
