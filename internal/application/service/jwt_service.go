package service

import (
	"fmt"
	"mcanvr/example-golang-api-with-fiber/pkg/constants"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

// JWTService handles token generation and validation
type JWTService struct {
	secretKey     string
	expirationHrs int
}

// NewJWTService creates a new JWT service instance
func NewJWTService(secretKey string, expirationHrs int) *JWTService {
	return &JWTService{
		secretKey:     secretKey,
		expirationHrs: expirationHrs,
	}
}

// GenerateToken creates a new JWT token for a user
func (s *JWTService) GenerateToken(userID int, username string, isAdmin bool) (string, error) {
	claims := jwt.MapClaims{
		"user_id":  userID,
		"username": username,
		"admin":    isAdmin,
		"exp":      time.Now().Add(time.Hour * time.Duration(s.expirationHrs)).Unix(),
		"iat":      time.Now().Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", fmt.Errorf("%s: %w", constants.TokenCreationFailed, err)
	}

	return tokenString, nil
}

// ValidateToken verifies the validity of a token and returns its claims
func (s *JWTService) ValidateToken(tokenString string) (*jwt.Token, jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validate the algorithm
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("%s: %v", constants.TokenInvalid, token.Header["alg"])
		}
		return []byte(s.secretKey), nil
	})

	if err != nil {
		if err.Error() == "Token is expired" {
			return nil, nil, fmt.Errorf(constants.TokenExpired)
		}
		return nil, nil, fmt.Errorf("%s: %w", constants.TokenInvalid, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		return nil, nil, fmt.Errorf(constants.TokenInvalid)
	}

	return token, claims, nil
}

// GetSecretKey returns the secret key
func (s *JWTService) GetSecretKey() string {
	return s.secretKey
}
