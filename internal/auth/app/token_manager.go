package auth

import (
	"context"
	"errors"
	context_keys "genpasstore/internal"
	httpx "genpasstore/internal/httpx/handler"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type TokenManager struct {
	secret []byte
}

func NewTokenManager(secret string, ttl time.Duration) *TokenManager {
	return &TokenManager{
		secret: []byte(secret),
	}
}

func (tokenManager *TokenManager) GenerateAccessToken(userID string, ttl time.Duration) (string, error) {
	return tokenManager.GenerateToken(userID, "access", ttl)
}

func (tokenManager *TokenManager) GenerateRefreshToken(userID string, ttl time.Duration) (string, error) {
	return tokenManager.GenerateToken(userID, "refresh", ttl)
}

func (tokenManger *TokenManager) GenerateToken(userID, typeToken string, ttl time.Duration) (string, error) {
	now := time.Now()
	exp := now.Add(ttl)

	claims := jwt.MapClaims{
		"iat":  now.Unix(),
		"exp":  exp.Unix(),
		"sub":  userID,
		"type": typeToken,
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

func (tokenManager *TokenManager) MiddlewareJWTToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		auth := r.Header.Get("Authorization")
		if !strings.HasPrefix(auth, "Bearer") {
			httpx.WriteError(w, http.StatusUnauthorized, "Invalid header request", httpx.ErrorDetails{
				"err": errors.New("invalid header for authorization, check documentation"),
			})
			return
		}
		tokenStr := strings.TrimPrefix(auth, "Bearer")
		tokenStr = strings.TrimSpace(tokenStr)

		userId, err := tokenManager.ValidateToken(tokenStr)
		if err != nil {
			httpx.WriteError(w, http.StatusUnauthorized, "Invalid token", httpx.ErrorDetails{
				"err": "Invalid access token. Mb need refresh?",
			})
			return
		}

		ctx := context.WithValue(r.Context(), context_keys.UserIDKey, userId)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
