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
	Index style
*/
func StylesIndex(c *fiber.Ctx) error {
	log.Println("get all styles")

	var styles []models.Style
	db.DB.Find(&styles)

	return c.JSON(fiber.Map{
		"data": styles,
	})
}

/*
	Create style
*/
func StylesCreate(c *fiber.Ctx) error {

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed create style: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed create style: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed create style: you need admin or developer permission",
		})
	}

	log.Println("start to create style")

	var style models.Style

	err := c.BodyParser(&style)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return err
	}
	db.DB.Create(&style)
	log.Printf("finish create style: %v", style.Style)

	return c.JSON(style)
}

/*
	UserShow
*/
func StyleShow(c *fiber.Ctx) error {
	style := controllerlogics.GetStyleFromId(c)

	//check style
	err := db.DB.First(&style).Error
	if err != nil {
		log.Printf("failed show style: style not found: id = %v", style.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed show style: style not found: id = %v", style.Id),
		})
	}

	log.Printf("start show style: id = %v", style.Id)

	db.DB.Find(&style)
	log.Printf("show user: id = %v, style = %v", style.Id, style.Style)

	return c.JSON(style)
}

/*
	UserUpdate
*/
func StyleUpdate(c *fiber.Ctx) error {
	style := controllerlogics.GetStyleFromId(c)

	//check account
	err := db.DB.First(&style).Error
	if err != nil {
		log.Printf("failed update style: style not found: id = %v", style.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed update style: style not found: id = %v", style.Id),
		})
	}

	log.Printf("start update style: id = %v", style.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed update style: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed update style: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed update style: you need admin or developer permission",
		})
	}

	err2 := c.BodyParser(&style)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return err2
	}

	db.DB.Model(&style).Updates(style)
	log.Println("success update style")

	return c.JSON(style)
}

/*
	UserDelete
*/
func StyleDelete(c *fiber.Ctx) error {
	style := controllerlogics.GetStyleFromId(c)

	//check account
	err := db.DB.First(&style).Error
	if err != nil {
		log.Printf("failed delete style: style not found: id = %v", style.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed delete style: style not found: id = %v", style.Id),
		})
	}

	log.Printf("start delete style: id = %v", style.Id)

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
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed delete style: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed delete style: you need admin or developer permission",
		})
	}

	db.DB.Delete(style)
	log.Println("success delete style")

	return c.JSON(fiber.Map{
		"message": "success delete style",
	})
}
