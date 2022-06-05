package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetHowtoFromId(c *fiber.Ctx) models.Howto {
	id, _ := strconv.Atoi(c.Params("id"))
	howto := models.Howto{
		Id: uint(id),
	}

	return howto
}
