package middleware

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"github.com/pkg/errors"
	"time"
)

func Token(userId string, hmacSecret []byte) (*string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 1).Unix(),
	})
	signedToken, err := token.SignedString(hmacSecret)
	if err != nil {
		return nil, err
	}
	return &signedToken, nil
}

func PareToken(tokenString string, hmacSecret []byte) (*string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return hmacSecret, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, errors.New("token is invalid")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("token claims is invalid")
	}
	id, ok := claims["userId"].(string)
	if !ok {
		return nil, errors.New("token claims is invalid")
	}
	return &id, nil
}
