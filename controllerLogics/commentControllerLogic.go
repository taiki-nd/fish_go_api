package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetCommentReplyFromId(c *fiber.Ctx) models.CommentReply {
	id, _ := strconv.Atoi(c.Params("id"))
	commentReply := models.CommentReply{
		Id: uint(id),
	}

	return commentReply
}
