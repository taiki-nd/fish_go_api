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
	Index Howto
*/
func HowtosIndex(c *fiber.Ctx) error {
	log.Println("get all howtos")

	var howtos []models.Howto
	db.DB.Find(&howtos)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   howtos,
	})
}

/*
	Create Howto
*/
func HowtosCreate(c *fiber.Ctx) error {

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed create howto: please login")
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed create howto: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "failed create howto: you need admin or developer permission",
		})
	}

	log.Println("start to create howto")

	var howto models.Howto

	err := c.BodyParser(&howto)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("POST method error: %v", err),
		})
	}
	db.DB.Create(&howto)
	log.Printf("finish create howto: %v", howto.Howto)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   howto,
	})
}

/*
	Show Howto
*/
func HowtoShow(c *fiber.Ctx) error {
	howto := controllerlogics.GetHowtoFromId(c)

	//check howto
	err := db.DB.First(&howto).Error
	if err != nil {
		log.Printf("failed show howto: howto not found: id = %v", howto.Id)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("failed show howto: howto not found: id = %v", howto.Id),
		})
	}

	log.Printf("start show howto: id = %v", howto.Id)

	db.DB.Find(&howto)
	log.Printf("show user: id = %v, howto = %v", howto.Id, howto.Howto)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   howto,
	})
}

/*
	Update Howto
*/
func HowtoUpdate(c *fiber.Ctx) error {
	howto := controllerlogics.GetHowtoFromId(c)

	//check account
	err := db.DB.First(&howto).Error
	if err != nil {
		log.Printf("failed update howto: howto not found: id = %v", howto.Id)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("failed update howto: howto not found: id = %v", howto.Id),
		})
	}

	log.Printf("start update howto: id = %v", howto.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed update howto: please login")
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed update howto: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "failed update howto: you need admin or developer permission",
		})
	}

	err2 := c.BodyParser(&howto)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("put method error: %v", err2),
		})
	}

	db.DB.Model(&howto).Updates(howto)
	log.Println("success update howto")

	return c.JSON(fiber.Map{
		"status": true,
		"data":   howto,
	})
}

/*
	Delete Howto
*/
func HowtoDelete(c *fiber.Ctx) error {
	howto := controllerlogics.GetHowtoFromId(c)

	//check account
	err := db.DB.First(&howto).Error
	if err != nil {
		log.Printf("failed delete howto: howto not found: id = %v", howto.Id)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("failed delete howto: howto not found: id = %v", howto.Id),
		})
	}

	log.Printf("start delete howto: id = %v", howto.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed delete user: please login")
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed delete howto: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "failed delete howto: you need admin or developer permission",
		})
	}

	db.DB.Delete(howto)
	log.Println("success delete howto")

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success delete howto",
	})
}
