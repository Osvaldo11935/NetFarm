package common

import "github.com/dgrijalva/jwt-go"

type Claims struct {
	Email  string   `json:"email"`
	UserID string   `json:"user_id"`
	Claims []string `json:"claims"`
	jwt.StandardClaims
}
