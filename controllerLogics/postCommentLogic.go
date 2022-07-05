package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetPostCommentFromId(c *fiber.Ctx) models.PostComment {
	id, _ := strconv.Atoi(c.Params("id"))
	postComment := models.PostComment{
		Id: uint(id),
	}

	return postComment
}
