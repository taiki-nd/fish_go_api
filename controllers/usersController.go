package controllers

import (
	"fish_go_api/db"
	"fish_go_api/models"
	"log"

	"github.com/gofiber/fiber/v2"
)

func UsersIndex(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}

func UsersCreate(c *fiber.Ctx) error {
	log.Println("start to create user")

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return err
	}

	if data["password"] != data["password_confirm"] {
		c.Status(400)
		log.Println("password & password_confirm dose not match.")
		return c.JSON(fiber.Map{
			"message": "password & password_confirm dose not match.",
		})
	}

	user := models.User{
		FirstName:      data["first_name"],
		LastName:       data["last_name"],
		Email:          data["email"],
		PermissionType: data["permission_type"],
	}
	user.SetPassword(data["password"])

	db.DB.Create(&user)
	log.Printf("finish register: %v %v", user.FirstName, user.LastName)

	return c.JSON(user)
}
