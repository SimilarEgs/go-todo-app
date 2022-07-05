package utils

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)

// payload data of the token
type TokenClaims struct {
	ID       uuid.UUID `json:"id"`
	Username string    `json:"username"`
	jwt.StandardClaims
}

// this function return new JWT
func CreateToken(username string, duration time.Duration) (string, error) {

	// setting JWT secure key
	key := os.Getenv("JWT_SECURITY_KEY")
	var jwtKey = []byte(key)

	// seting tokenID
	tokenId, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}

	// createing JWT claims
	claims := &TokenClaims{
		ID:       tokenId,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(duration).UnixNano(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	// creating a new token with declared claims
	Jwttoken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := Jwttoken.SignedString([]byte(jwtKey))
	return ("Bearer " + tokenString), err

}

func VerifyJWT(tokenString string) (*TokenClaims, error) {

	jwtKey := []byte(os.Getenv("JWT_SECURITY_KEY"))

	claims := &TokenClaims{}

	token, err := jwt.ParseWithClaims(tokenString, claims, func(t *jwt.Token) (interface{}, error) {
		// verifying token algorithm
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected siging method")
		}
		return []byte(jwtKey), nil
	})

	if !token.Valid {
		return nil, err
	}

	return claims, nil
}
