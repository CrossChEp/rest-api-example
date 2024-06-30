package models

import "github.com/golang-jwt/jwt"

type Claims struct {
	UserID UserID `json:"userID"`
	Email  string `json:"email"`

	jwt.StandardClaims
}
