package controllers

import (
	"fish_go_api/controllerlogics"
	"fish_go_api/db"
	"fish_go_api/models"
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

/*
	Index post
*/
func PostsIndex(c *fiber.Ctx) error {
	log.Println("get all posts")

	var posts []models.Post
	db.DB.Preload("PostComments").Find(&posts)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   posts,
	})
}

/*
	Create post
*/
func PostsCreate(c *fiber.Ctx) error {

	log.Println("start to create post")

	var post models.Post

	err := c.BodyParser(&post)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("POST method error: %v", err),
		})
	}

	db.DB.Create(&post)
	log.Printf("finish create post: %v", post.Name)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   post,
	})
}

/*
	Show Post
*/
func PostShow(c *fiber.Ctx) error {
	post := controllerlogics.GetPostFromId(c)

	//check post
	err := db.DB.First(&post).Error
	if err != nil {
		log.Printf("failed show post: post not found: id = %v", post.Id)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("failed show post: post not found: id = %v", post.Id),
		})
	}

	log.Printf("start show post: id = %v", post.Id)

	db.DB.Preload("PostComments").Find(&post)
	log.Printf("show post: id = %v, Name = %v", post.Id, post.Name)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   post,
	})
}

/*
	Update post
*/
func PostUpdate(c *fiber.Ctx) error {
	post := controllerlogics.GetPostFromId(c)

	err2 := c.BodyParser(&post)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("put method error: %v", err2),
		})
	}

	db.DB.Model(&post).Updates(post)
	log.Println("success update post")

	return c.JSON(fiber.Map{
		"status": true,
		"data":   post,
	})
}

/*
	Delete post
*/
func PostDelete(c *fiber.Ctx) error {
	post := controllerlogics.GetPostFromId(c)

	// delete asociation (transaction)
	err := db.DB.Transaction(func(tx *gorm.DB) error {
		// var postComment_id []int64
		// tx.Table("post_comments").Where("post_id = ?", post.Id).Pluck("id", &postComment_id)
		// if len(postComment_id) != 0 {
		// 	err := tx.Table("comment_replies").Where("post_comment_id IN (?)", postComment_id).Delete("").Error
		// 	if err != nil {
		// 		return err
		// 	}
		// }

		err := tx.Table("post_comments").Where("post_id = ?", post.Id).Delete("").Error
		if err != nil {
			return err
		}

		return nil
	})
	if err != nil {
		log.Printf("transaction error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("transaction error: %v", err),
		})
	}

	// if post.Filename != "" {
	// 	err := ImageDelete(post.Filename)
	// 	if err != "" {
	// 		log.Println(err)
	// 		c.JSON(fiber.Map{
	// 			"status":  false,
	// 			"message": err,
	// 		})
	// 	}
	// }

	db.DB.Delete(post)
	log.Println("success delete post")

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success delete post",
	})
}
