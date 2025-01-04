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
var appEnv = os.Getenv("APP_ENV")

type Claims struct {
	Id       uint64 `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
	ClassId   uint64 `json:"classId"`
	jwt.RegisteredClaims
}

func GenerateAndSetAccessToken(w http.ResponseWriter, id uint64, username string, email string, role string, classId uint64) (string, error) {
	Exp := time.Now().Add(24 * time.Hour * 7)

	claims := Claims{
		Id:       id,
		Username: username,
		Email:    email,
		Role:     role,
		ClassId: classId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(Exp),
		},
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	accessTokenStr, err := accessToken.SignedString(accessTokenSecret)

	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "AccessToken",
		Value:    accessTokenStr,
		Expires:  Exp,
		Secure:   appEnv == "production",
		HttpOnly: true,
		Path:     "/",
	})

	return accessTokenStr, nil
}

func GenerateAndSetRefreshToken(w http.ResponseWriter, id uint64, username string, email string, role string, classId uint64) (string, error) {
	Exp := time.Now().Add(24 * time.Hour * 7)

	claims := Claims{
		Id:       id,
		Username: username,
		Email:    email,
		Role:     role,
		ClassId: classId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(Exp),
		},
	}

	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	refreshTokenStr, err := refreshToken.SignedString(refreshTokenSecret)

	if err != nil {
		return "", err
	}

	http.SetCookie(w, &http.Cookie{
		Name:     "RefreshToken",
		Value:    refreshTokenStr,
		Expires:  Exp,
		Secure:   appEnv == "production",
		HttpOnly: true,
		Path:     "/",
	})

	return refreshTokenStr, nil
}

func DecodeRefreshToken(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie("RefreshToken")

	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	return claims, nil
}

func DecodeAccessToken(r *http.Request) (*Claims, error) {
	cookie, err := r.Cookie("AccessToken")

	if err != nil {
		return nil, err
	}

	token, err := jwt.ParseWithClaims(cookie.Value, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})

	if err != nil || !token.Valid {
		return nil, err
	}

	claims, ok := token.Claims.(*Claims)

	if !ok {
		return nil, fmt.Errorf("failed to parse claims")
	}

	return claims, nil
}

func ValidateAccessToken(accessToken string) error {
	token, err := jwt.ParseWithClaims(accessToken, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return accessTokenSecret, nil
	})

	if err != nil || !token.Valid {
		return err
	}

	_, ok := token.Claims.(*Claims)

	if !ok {
		return fmt.Errorf("failed to parse claims")
	}

	return nil
}
