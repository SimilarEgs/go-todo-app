package service

import (
	"log"
	"os"
	"time"

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

	// parsing time duration for JWT
	duration, err := time.ParseDuration(os.Getenv("JWT_TOKEN_DURATION"))
	if err != nil {
		return "", err
	}

	// genereate token
	token, err := utils.CreateToken(username, duration)
	if err != nil {
		return "", err
	}
	log.Println(token)

	return token, err
}
