package util

import (
	"errors"
	"eventsie/auth/config"
	"fmt"

	"github.com/golang-jwt/jwt/v4"
)

type JWTData struct {
	ID        string
	FirstName string
	LastName  string
	Email     string
}

func GenerateJWTToken(id string, firstName string, lastName string, email string) (string, error) {
	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":        id,
		"firstName": firstName,
		"lastName":  lastName,
		"email":     email,
	})

	cfg := config.GetConfig()
	tokenString, err := token.SignedString([]byte(cfg.JWT_SECRET))

	if err != nil {
		return "", errors.New("could not sign jwt token")
	}

	return tokenString, nil
}

func ParseJWTToken(tokenString string) (*JWTData, error) {
	cfg := config.GetConfig()

	// Parse JWT token
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.JWT_SECRET), nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return &JWTData{
			ID:        claims["id"].(string),
			FirstName: claims["firstName"].(string),
			LastName:  claims["lastName"].(string),
			Email:     claims["email"].(string),
		}, nil
	} else {
		return nil, fmt.Errorf("could not verify token")
	}
}
