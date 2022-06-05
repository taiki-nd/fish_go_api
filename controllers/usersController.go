package controllers

import (
	"fish_go_api/controllerlogics"
	"fish_go_api/db"
	"fish_go_api/models"
	"fmt"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
)

/*
	Index user
*/
func UsersIndex(c *fiber.Ctx) error {
	log.Println("get all users")
	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed show all user: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	var users []models.User
	db.DB.Find(&users)

	return c.JSON(fiber.Map{
		"data": users,
	})
}

/*
	Create user
*/
func UsersCreate(c *fiber.Ctx) error {
	log.Println("start to create user")

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed create user: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" {
		log.Println("failed create user: you need admin permission")
		return c.JSON(fiber.Map{
			"message": "failed create user: you need admin permission",
		})
	}

	var data map[string]string

	err := c.BodyParser(&data)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return err
	}

	// check email unique
	var createdUser models.User
	db.DB.Where("email = ?", data["email"]).First(&createdUser)
	if createdUser.Id != 0 {
		log.Printf("failed create user: exists account: mail = %v", data["email"])
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed create user: exists account: mail = %v", data["email"]),
		})
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
	log.Printf("finish create user: %v %v", user.FirstName, user.LastName)

	return c.JSON(user)
}

/*
	Show User
*/
func UserShow(c *fiber.Ctx) error {
	user := controllerlogics.GetUserFromId(c)

	//check account
	err := db.DB.First(&user).Error
	if err != nil {
		log.Printf("failed show user: user not found: id = %v", user.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed show user: user not found: id = %v", user.Id),
		})
	}

	log.Printf("start show user: id = %v", user.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed show user: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	db.DB.Find(&user)
	log.Printf("show user: id = %v", user.Id)

	return c.JSON(user)
}

/*
	Update User
*/
func UserUpdate(c *fiber.Ctx) error {
	user := controllerlogics.GetUserFromId(c)

	//check account
	err := db.DB.First(&user).Error
	if err != nil {
		log.Printf("failed update user: user not found: id = %v", user.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed update user: user not found: id = %v", user.Id),
		})
	}

	log.Printf("start update user: id = %v", user.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed update user: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" {
		log.Println("failed update user: you need admin permission")
		return c.JSON(fiber.Map{
			"message": "failed update user: you need admin permission",
		})
	}

	err2 := c.BodyParser(&user)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return err2
	}

	db.DB.Model(&user).Updates(user)
	log.Println("success update user")

	return c.JSON(user)
}

/*
	Delete User
*/
func UserDelete(c *fiber.Ctx) error {
	user := controllerlogics.GetUserFromId(c)

	//check account
	err := db.DB.First(&user).Error
	if err != nil {
		log.Printf("failed delete user: user not found: id = %v", user.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed delete user: user not found: id = %v", user.Id),
		})
	}

	log.Printf("start delete user: id = %v", user.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed delete user: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" {
		log.Println("failed delete user: you need admin permission")
		return c.JSON(fiber.Map{
			"message": "failed delete user: you need admin permission",
		})
	}

	db.DB.Delete(user)
	log.Println("success delete user")

	return c.JSON(fiber.Map{
		"message": "success delete user",
	})
}

/*
	Login
*/
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

	log.Printf("success login: %v", data["email"])

	return c.JSON(token)
}

/*
	User infomation
*/
func User(c *fiber.Ctx) error {
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)

	var user models.User
	db.DB.Where("id =?", issuer).First(&user)

	return c.JSON(&user)
}

/*
	Logout
*/
func Logout(c *fiber.Ctx) error {
	issuer, _ := controllerlogics.ParseJwt(c.Cookies("jwt"))

	var user models.User
	db.DB.Where("id = ?", issuer).First(&user)

	log.Printf("start logout: id = %v", user.Id)

	cookie := fiber.Cookie{
		Name:     "jwt",
		Value:    "",
		Expires:  time.Now().Add(-time.Hour),
		HTTPOnly: true,
	}
	c.Cookie(&cookie)
	log.Println("success logout")

	return c.JSON(fiber.Map{
		"message": "success logout",
	})
}
