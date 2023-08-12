package util

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte("tiktok")

type Claims struct {
	Username string `json:"username"`
	Password string `json:"password"`
	Id       uint   `json:"user_id"`
	jwt.RegisteredClaims
}

func GenerateToken(username, password string, id uint) (string, error) {

	claims := Claims{
		username,
		password,
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * 24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		}}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func ParseToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	claims, ok := token.Claims.(*Claims)
	if ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
