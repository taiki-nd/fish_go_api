package controllerlogics

import (
	"fish_go_api/models"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go/v4"
	"github.com/gofiber/fiber/v2"
)

type Claims struct {
	jwt.StandardClaims
}

func GetUserFromId(c *fiber.Ctx) models.User {
	id, _ := strconv.Atoi(c.Params("id"))
	user := models.User{
		Id: uint(id),
	}

	return user
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
