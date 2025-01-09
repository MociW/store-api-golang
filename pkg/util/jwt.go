package util

import (
	"time"

	"github.com/MociW/store-api-golang/config"
	"github.com/MociW/store-api-golang/internal/user/model"
	"github.com/golang-jwt/jwt/v5"
)

type Claim struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateToken(user *model.User, cfg *config.Config, expire time.Duration) (string, error) {
	claims := Claim{
		ID:    user.UserID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expire)),
			Issuer:    "jwt",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte(cfg.Server.JWTSecretKey))
	if err != nil {
		return "", err
	}

	return t, nil
}

func GenerateTokenPair(user *model.User, cfg *config.Config) (accToken string, refToken string, err error) {
	accToken, err = GenerateToken(user, cfg, 15*time.Minute) // 15 minute
	if err != nil {
		return
	}
	refToken, err = GenerateToken(user, cfg, 1*24*time.Hour) // 1 day
	return
}

func ValidateJwtToken(tokenString string, cfg *config.Config) (*Claim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claim{}, func(token *jwt.Token) (interface{}, error) {
		secretKey := []byte(cfg.Server.JWTSecretKey)
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*Claim)
	if !ok || !token.Valid {
		return nil, err
	}

	return claims, nil
}
