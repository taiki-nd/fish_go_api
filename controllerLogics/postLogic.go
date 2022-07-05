package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPostFromId(c *fiber.Ctx) models.Post {
	id, _ := strconv.Atoi(c.Params("id"))
	post := models.Post{
		Id: uint(id),
	}

	return post
}
