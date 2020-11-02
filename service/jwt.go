package service

import (
	"belajar-go-rest-api/config"
	"belajar-go-rest-api/entities"
	"errors"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// JWTService struct
type JWTService struct {
	Config *config.JWTConfig
}

// NewJWTService func
func NewJWTService() *JWTService {
	jwtConfig := config.ConfigureJWT()

	return &JWTService{
		Config: jwtConfig,
	}
}

// Create func
func (j JWTService) Create(userData *entities.User) (string, error) {
	claims := &entities.JWTClaims{}
	claims.User = userData
	claims.ExpiresAt = int64(time.Hour) * 3

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	t, err := token.SignedString([]byte(j.Config.Secret))
	if err != nil {
		return "", errors.New("Token generation fail")
	}

	return t, nil
}
