package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetGroundFromId(c *fiber.Ctx) models.Ground {
	id, _ := strconv.Atoi(c.Params("id"))
	ground := models.Ground{
		Id: uint(id),
	}

	return ground
}
