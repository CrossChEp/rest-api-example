package models

import (
	"github.com/golang-jwt/jwt"
	"time"
)

type Claims struct {
	UserID    UserID    `json:"userID"`
	Email     string    `json:"email"`
	ExpiresAt time.Time `json:"expiresAt"`

	jwt.StandardClaims
}
