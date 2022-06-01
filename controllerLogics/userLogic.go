package controllerlogics

import (
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
)

type Claims struct {
	jwt.StandardClaims
}

func GenerateJwt(userID int) (string, error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		Issuer:    strconv.Itoa(int(userID)),
		ExpiresAt: &jwt.Time{time.Now().Add(time.Hour * 24)},
	})

	return claims.SignedString([]byte("secret"))
}

func ParseJwt(cookie string) (string, error) {
	token, err := jwt.ParseWithClaims(cookie, &Claims{}, func(t *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil || !token.Valid {
		return "", nil
	}

	claims := token.Claims.(*Claims)

	return claims.Issuer, nil
}
