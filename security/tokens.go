package security

import (
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"strconv"
	"time"
)

const (
	AccessTokenExpire  = time.Hour * 15
	RefreshTokenExpire = time.Hour * 24 * 7
)

var (
	// Replace with your own secret key
	accessSecret  = []byte("your-access-secret-key")
	refreshSecret = []byte("your-refresh-secret-key")
)

func GenerateAccessToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpire)),
		ID:        strconv.Itoa(int(userID)),
	})

	signedToken, err := token.SignedString(accessSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func GenerateRefreshToken(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(AccessTokenExpire)),
		ID:        strconv.Itoa(int(userID)),
	})

	signedToken, err := token.SignedString(refreshSecret)
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func ValidateAccessToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return accessSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		userID, err := strconv.Atoi(claims.ID)
		if err != nil {
			return 0, fmt.Errorf("invalid access token: %v", err)
		}
		return uint(userID), nil
	}

	return 0, fmt.Errorf("invalid access token: %v", err)
}

func ValidateRefreshToken(tokenString string) (uint, error) {
	token, err := jwt.ParseWithClaims(tokenString, &jwt.RegisteredClaims{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecret, nil
	})

	if err != nil {
		return 0, err
	}

	if claims, ok := token.Claims.(*jwt.RegisteredClaims); ok && token.Valid {
		userID, err := strconv.Atoi(claims.ID)
		if err != nil {
			return 0, errors.New("invalid user ID in refresh token")
		}
		return uint(userID), nil
	}

	return 0, errors.New("invalid refresh token")
}
