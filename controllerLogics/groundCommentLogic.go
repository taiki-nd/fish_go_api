package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetGroundCommentFromId(c *fiber.Ctx) models.GroundComment {
	id, _ := strconv.Atoi(c.Params("id"))
	groundComment := models.GroundComment{
		Id: uint(id),
	}

	return groundComment
}
