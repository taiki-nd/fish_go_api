package controllers

import (
	"fish_go_api/controllerlogics"
	"fish_go_api/db"
	"fish_go_api/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

/*
	Index groundComment
*/
func GroundCommentsIndex(c *fiber.Ctx) error {
	log.Println("get all groundComments")

	var groundComments []models.GroundComment
	db.DB.Preload("CommentReplies").Find(&groundComments)

	return c.JSON(fiber.Map{
		"data": groundComments,
	})
}

/*
	Create groundComment
*/
func GroundCommentsCreate(c *fiber.Ctx) error {
	log.Println("start to create groundComment")

	var groundComment models.GroundComment

	err := c.BodyParser(&groundComment)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return err
	}
	db.DB.Create(&groundComment)
	log.Printf("finish create groundComment: %v", groundComment.Id)

	return c.JSON(groundComment)
}

/*
	Show groundComment
*/
func GroundCommentShow(c *fiber.Ctx) error {
	groundComment := controllerlogics.GetGroundCommentFromId(c)

	//check groundComment
	err := db.DB.First(&groundComment).Error
	if err != nil {
		log.Printf("failed show groundComment: groundComment not found: id = %v", groundComment.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed show groundComment: groundComment not found: id = %v", groundComment.Id),
		})
	}

	log.Printf("start show groundComment: id = %v", groundComment.Id)

	db.DB.Preload("CommentReplies").Find(&groundComment)
	log.Printf("show user: id = %v, groundComment = %v", groundComment.Id, groundComment.Id)

	return c.JSON(groundComment)
}

/*
	Update groundComment
*/
func GroundCommentUpdate(c *fiber.Ctx) error {
	groundComment := controllerlogics.GetGroundCommentFromId(c)

	log.Printf("start update groundComment: id = %v", groundComment.Id)

	err2 := c.BodyParser(&groundComment)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return err2
	}

	db.DB.Model(&groundComment).Updates(groundComment)
	log.Println("success update groundComment")

	return c.JSON(groundComment)
}

/*
	Delete groundComment
*/
func GroundCommentDelete(c *fiber.Ctx) error {
	groundComment := controllerlogics.GetGroundCommentFromId(c)

	log.Printf("start delete groundComment: id = %v", groundComment.Id)

	db.DB.Table("comment_replies").Where("ground_comment_id = ?", groundComment.Id).Delete("")
	db.DB.Delete(groundComment)
	log.Println("success delete groundComment")

	return c.JSON(fiber.Map{
		"message": "success delete groundComment",
	})
}
