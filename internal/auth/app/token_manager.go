package auth

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	secret []byte
	ttl    time.Duration
}

func NewTokenManager(secret string, ttl time.Duration) *TokenManager {
	return &TokenManager{
		secret: []byte(secret),
		ttl:    ttl,
	}
}

func (tokenManger *TokenManager) GenerateToken(userID string) (string, error) {
	now := time.Now()
	exp := now.Add(tokenManger.ttl)

	claims := jwt.MapClaims{
		"now": now.Unix(),
		"exp": exp.Unix(),
		"sub": userID,
	}

	keyWithClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedKeyWithClaims, err := keyWithClaims.SignedString(tokenManger.secret)
	if err != nil {
		return "", err
	}

	return signedKeyWithClaims, nil
}

func (tokenManager *TokenManager) ValidateToken(tokenStr string) (string, error) {
	if tokenStr == "" {
		return "", errors.New("token is empty")
	}

	claims := jwt.MapClaims{}

	token, err := jwt.ParseWithClaims(tokenStr, claims, func(t *jwt.Token) (any, error) {
		if t.Method != jwt.SigningMethodHS256 {
			return nil, errors.ErrUnsupported
		}
		return tokenManager.secret, nil
	})
	if err != nil {
		return "", err
	}

	if !token.Valid {
		return "", errors.New("invalid token")
	}

	sub, ok := claims["sub"].(string)
	if !ok || sub == "" {
		return "", errors.New("sub is empty")
	}
	return sub, nil
}
