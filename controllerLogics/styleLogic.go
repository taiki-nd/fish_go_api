package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetStyleFromId(c *fiber.Ctx) models.Style {
	id, _ := strconv.Atoi(c.Params("id"))
	style := models.Style{
		Id: uint(id),
	}

	return style
}
