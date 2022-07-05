package utils

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt"
)

// payload data of the token
type TokenClaims struct {
	ID       int64  `json:"userid"`
	Username string `json:"username"`
	jwt.StandardClaims
}

// this function return new JWT
func CreateToken(username string, userid int64) (string, error) {

	// setting JWT secure key
	key := os.Getenv("JWT_SECURITY_KEY")
	var jwtKey = []byte(key)

	// extracting duration time of the token from config
	durationStr := os.Getenv("JWT_TOKEN_DURATION_IN_MINUTE")
	duration, err := strconv.Atoi(durationStr)
	if err != nil {
		return "", err
	}

	// setting expiration time of the token
	expTokenTime := time.Now().Add(time.Duration(duration) * time.Minute).Unix()

	// createing JWT claims
	claims := &TokenClaims{
		ID:       userid,
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTokenTime,
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
