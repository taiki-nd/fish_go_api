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
	Index CommentReply
*/
func CommentRepliesIndex(c *fiber.Ctx) error {
	log.Println("get all commentReplies")

	var commentReplies []models.CommentReply
	db.DB.Find(&commentReplies)

	return c.JSON(fiber.Map{
		"data": commentReplies,
	})
}

/*
	Create commentReply
*/
func CommentRepliesCreate(c *fiber.Ctx) error {
	log.Println("start to create commentReply")

	var commentReply models.CommentReply

	err := c.BodyParser(&commentReply)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return err
	}
	db.DB.Create(&commentReply)
	log.Printf("finish create commentReply: %v", commentReply.Id)

	return c.JSON(commentReply)
}

/*
	Show CommentReply
*/
func CommentReplyShow(c *fiber.Ctx) error {
	commentReply := controllerlogics.GetCommentReplyFromId(c)

	//check commentReply
	err := db.DB.First(&commentReply).Error
	if err != nil {
		log.Printf("failed show commentReply: commentReply not found: id = %v", commentReply.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed show commentReply: commentReply not found: id = %v", commentReply.Id),
		})
	}

	log.Printf("start show commentReply: id = %v", commentReply.Id)

	db.DB.Find(&commentReply)
	log.Printf("show user: id = %v, commentReply = %v", commentReply.Id, commentReply.Id)

	return c.JSON(commentReply)
}

/*
	Update CommentReply
*/
func CommentReplyUpdate(c *fiber.Ctx) error {
	commentReply := controllerlogics.GetCommentReplyFromId(c)

	log.Printf("start update commentReply: id = %v", commentReply.Id)

	err2 := c.BodyParser(&commentReply)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return err2
	}

	db.DB.Model(&commentReply).Updates(commentReply)
	log.Println("success update commentReply")

	return c.JSON(commentReply)
}

/*
	Delete CommentReply
*/
func CommentReplyDelete(c *fiber.Ctx) error {
	commentReply := controllerlogics.GetCommentReplyFromId(c)

	log.Printf("start delete commentReply: id = %v", commentReply.Id)

	db.DB.Delete(commentReply)
	log.Println("success delete commentReply")

	return c.JSON(fiber.Map{
		"message": "success delete commentReply",
	})
}
