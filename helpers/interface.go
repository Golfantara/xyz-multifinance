package helpers

import (
	"github.com/golang-jwt/jwt/v5"
)

type JWTInterface interface {
	GenerateJWT(userID string, roleID string) map[string]any
	GenerateToken(userID string, roleID string) string
	ExtractToken(token *jwt.Token) any
	ValidateToken(token string, secret string) (*jwt.Token, error)
	RefereshJWT(refreshToken *jwt.Token) map[string]any
}