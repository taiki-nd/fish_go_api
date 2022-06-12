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
	log.Println("get all groundCommments")

	var groundCommments []models.GroundComment
	db.DB.Find(&groundCommments)

	return c.JSON(fiber.Map{
		"data": groundCommments,
	})
}

/*
	Create groundCommment
*/
func GroundCommentsCreate(c *fiber.Ctx) error {
	log.Println("start to create groundCommment")

	var groundCommment models.GroundComment

	err := c.BodyParser(&groundCommment)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return err
	}
	db.DB.Create(&groundCommment)
	log.Printf("finish create groundCommment: %v", groundCommment.Id)

	return c.JSON(groundCommment)
}

/*
	Show groundComment
*/
func GroundCommentShow(c *fiber.Ctx) error {
	groundCommment := controllerlogics.GetGroundCommentFromId(c)

	//check groundCommment
	err := db.DB.First(&groundCommment).Error
	if err != nil {
		log.Printf("failed show groundCommment: groundCommment not found: id = %v", groundCommment.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed show groundCommment: groundCommment not found: id = %v", groundCommment.Id),
		})
	}

	log.Printf("start show groundCommment: id = %v", groundCommment.Id)

	db.DB.Find(&groundCommment)
	log.Printf("show user: id = %v, groundCommment = %v", groundCommment.Id, groundCommment.Id)

	return c.JSON(groundCommment)
}

/*
	Update groundComment
*/
func GroundCommentUpdate(c *fiber.Ctx) error {
	groundCommment := controllerlogics.GetGroundCommentFromId(c)

	log.Printf("start update groundCommment: id = %v", groundCommment.Id)

	err2 := c.BodyParser(&groundCommment)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return err2
	}

	db.DB.Model(&groundCommment).Updates(groundCommment)
	log.Println("success update groundCommment")

	return c.JSON(groundCommment)
}

/*
	Delete groundComment
*/
func GroundCommentDelete(c *fiber.Ctx) error {
	groundCommment := controllerlogics.GetGroundCommentFromId(c)

	log.Printf("start delete groundCommment: id = %v", groundCommment.Id)

	db.DB.Delete(groundCommment)
	log.Println("success delete groundCommment")

	return c.JSON(fiber.Map{
		"message": "success delete groundCommment",
	})
}
