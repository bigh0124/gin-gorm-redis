package utils

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func PasswordMatches(plaintextPassword string, hashedPassword string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(plaintextPassword))
	if err != nil {
		return false, err
	}

	return true, nil
}

func GenerateJWT(username string) (string, error) {
	now := time.Now()
	claims := jwt.RegisteredClaims{
		Subject:   username,
		ExpiresAt: jwt.NewNumericDate(now.Add(time.Hour * 72)),
		IssuedAt:  jwt.NewNumericDate(now),
		NotBefore: jwt.NewNumericDate(now),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte("secret"))
}

func ParseJWT(tokenValue string) (string, error) {
	parts := strings.Split(tokenValue, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return "", errors.New("invalid token format: expected 'Bearer <token>'")
	}

	token, err := jwt.ParseWithClaims(parts[1], &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})

	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}

	claims, ok := token.Claims.(*jwt.RegisteredClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	return claims.Subject, nil
}
