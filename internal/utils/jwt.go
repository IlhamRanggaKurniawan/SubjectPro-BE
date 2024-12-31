package utils

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenType string

const (
	AccessToken  TokenType = "Access Token"
	RefreshToken TokenType = "Refresh Token"
)

var refreshTokenSecret = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))
var accessTokenSecret = []byte(os.Getenv("ACCESS_TOKEN_SECRET"))

type Claims struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	Motto    string `json:"motto"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(id uint64, username string, email string, role string, motto *string) (string, error) {
	Exp := time.Now().Add(15 * time.Minute)

	claims := Claims{
		Id:       id,
		Username: username,
		Email:    email,
		Role:     role,
		Motto:    *motto,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(Exp),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenStr, err := accessToken.SignedString(accessTokenSecret)

	if err != nil {
		return "", err
	}

	return accessTokenStr, nil
}

func GenerateRefreshToken(id uint64, username string, email string, role string, motto *string) (string, error) {
	Exp := time.Now().Add(24 * time.Hour * 7)

	claims := Claims{
		Id:       id,
		Username: username,
		Email:    email,
		Role:     role,
		Motto:    *motto,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(Exp),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshTokenStr, err := refreshToken.SignedString(refreshTokenSecret)

	if err != nil {
		return "", err
	}

	return refreshTokenStr, nil
}

func ValidateToken(tokenString string, tokenType TokenType) (*Claims, error) {
	Claims := &Claims{}
	var secretKey []byte

	if tokenType == AccessToken {
		secretKey = accessTokenSecret
	} else {
		secretKey = refreshTokenSecret
	}

	token, err := jwt.ParseWithClaims(tokenString, Claims, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, fmt.Errorf("invalid token")
	}

	if Claims.ExpiresAt.Before(time.Now()) {
		return nil, fmt.Errorf("token expired")
	}

	return Claims, nil
}

func DecodeAccessToken(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie("AccessToken")

	if err != nil {
		return nil, fmt.Errorf("could not find access token in cookies: %v", err)
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	claims, ok := token.Claims.(*Claims)
	
	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	return claims, nil
}