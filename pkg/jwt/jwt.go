// Package jwt is a permission authentication processing middleware
package jwt

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"time"
)

type TokenInfo struct {
	Id int64 `json:"id"`
	jwt.StandardClaims
}

const (
	tokenValidDuration = 2 * time.Hour
	issuer             = "帅过吴彦组"
	secret             = "tiktok"
)

func GenToken(id int64) (string, error) {
	c := TokenInfo{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenValidDuration).Unix(),
			Issuer:    issuer,
		},
	}
	// Specifying a signature algorithm
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
	// Return token encoding
	return token.SignedString([]byte(secret))
}

func ParseToken(tokenString string) (*TokenInfo, error) {
	token, err := jwt.ParseWithClaims(tokenString, &TokenInfo{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	tokenInfo, ok := token.Claims.(*TokenInfo)
	if ok && token.Valid {
		return tokenInfo, nil
	}
	return nil, errors.New("invalid token")
}
