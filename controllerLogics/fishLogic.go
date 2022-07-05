package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetFishFromId(c *fiber.Ctx) models.Fish {
	id, _ := strconv.Atoi(c.Params("id"))
	fish := models.Fish{
		Id: uint(id),
	}

	return fish
}
