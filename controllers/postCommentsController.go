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
	Index postComment
*/
func PostCommentsIndex(c *fiber.Ctx) error {
	log.Println("get all postComments")

	var postComments []models.PostComment
	db.DB.Find(&postComments)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   postComments,
	})
}

/*
	Create postComment
*/
func PostCommentsCreate(c *fiber.Ctx) error {
	log.Println("start to create postComment")

	var postComment models.PostComment

	err := c.BodyParser(&postComment)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("POST method error: %v", err)})
	}
	db.DB.Create(&postComment)
	log.Printf("finish create postComment: %v", postComment.Id)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   postComment,
	})
}

/*
	Show postComment
*/
func PostCommentShow(c *fiber.Ctx) error {
	postComment := controllerlogics.GetPostCommentFromId(c)

	//check postComment
	err := db.DB.First(&postComment).Error
	if err != nil {
		log.Printf("failed show postComment: postComment not found: id = %v", postComment.Id)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("failed show postComment: postComment not found: id = %v", postComment.Id),
		})
	}

	log.Printf("start show postComment: id = %v", postComment.Id)

	db.DB.Find(&postComment)
	log.Printf("show user: id = %v, postComment = %v", postComment.Id, postComment.Id)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   postComment,
	})
}

/*
	Update postComment
*/
func PostCommentUpdate(c *fiber.Ctx) error {
	postComment := controllerlogics.GetPostCommentFromId(c)

	log.Printf("start update postComment: id = %v", postComment.Id)

	err2 := c.BodyParser(&postComment)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("put method error: %v", err2)})
	}

	db.DB.Model(&postComment).Updates(postComment)
	log.Println("success update postComment")

	return c.JSON(fiber.Map{
		"status": true,
		"data":   postComment,
	})
}

/*
	Delete postComment
*/
func PostCommentDelete(c *fiber.Ctx) error {
	postComment := controllerlogics.GetPostCommentFromId(c)

	//check record
	err := db.DB.First(&postComment).Error
	if err != nil {
		log.Printf("failed delete postComment: postComment not found: id = %v", postComment.Id)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("failed delete postComment: postComment not found: id = %v", postComment.Id),
		})
	}

	log.Printf("start delete postComment: id = %v", postComment.Id)
	log.Println(postComment.Filename)

	if postComment.Filename != "" {
		err := ImageDelete(postComment.Filename)
		if err != "" {
			log.Println(err)
			c.JSON(fiber.Map{
				"status":  false,
				"message": err,
			})
		}
	}

	//db.DB.Table("comment_replies").Where("post_comment_id = ?", postComment.Id).Delete("")
	db.DB.Delete(postComment)
	log.Println("success delete postComment")

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success delete postComment",
	})
}
