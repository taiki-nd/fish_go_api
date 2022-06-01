package controllers

import (
	"fish_go_api/controllerlogics"
	"fish_go_api/db"
	"fish_go_api/models"
	"log"
	"time"

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

func Login(c *fiber.Ctx) error {
	log.Println("start login")

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Printf("login error: %v", err)
		return err
	}

	var user models.User
	db.DB.Where("email = ?", data["email"]).First(&user)
	if user.Id == 0 {
		c.SendStatus(400)
		log.Printf("user not found: %v", data["email"])
		return c.JSON(fiber.Map{
			"message": "user not found",
		})
	}

	err = user.ComparePassword(data["password"])
	if err != nil {
		c.SendStatus(400)
		log.Printf("incorrect password: ID = %v", user.Id)
		return c.JSON(fiber.Map{
			"message": "incorrect password",
		})
	}

	token, err := controllerlogics.GenerateJwt(int(user.Id))
	if err != nil {
		return c.SendStatus(500)
	}

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    token,
		Expires:  time.Now().Add(time.Hour * 24),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)

	log.Printf("login success: %v", data["email"])

	return c.JSON(token)
}
