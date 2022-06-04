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
	Index ground
*/
func GroundsIndex(c *fiber.Ctx) error {
	log.Println("get all grounds")

	var grounds []models.Ground
	db.DB.Find(&grounds)

	return c.JSON(fiber.Map{
		"data": grounds,
	})
}

/*
	Create ground
*/
func GroundsCreate(c *fiber.Ctx) error {

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed create ground: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed create ground: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed create ground: you need admin or developer permission",
		})
	}

	log.Println("start to create ground")

	var groundStyle models.GroundAssociation

	err := c.BodyParser(&groundStyle)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return err
	}

	styles := controllerlogics.GetStyles(groundStyle.Styles)

	ground := models.Ground{
		Name:    groundStyle.Name,
		Address: groundStyle.Address,
		Tell:    groundStyle.Tell,
		Email:   groundStyle.Email,
		Break:   groundStyle.Break,
		Styles:  styles,
		Price:   groundStyle.Price,
		Url:     groundStyle.Url,
		Feature: groundStyle.Feature,
		Rule:    groundStyle.Rule,
		Other:   groundStyle.Other,
	}

	db.DB.Create(&ground)
	log.Printf("finish create ground: %v", ground.Name)

	return c.JSON(ground)
}

/*
	UserShow
*/
func GroundShow(c *fiber.Ctx) error {
	ground := controllerlogics.GetGroundFromId(c)

	//check ground
	err := db.DB.First(&ground).Error
	if err != nil {
		log.Printf("failed show ground: ground not found: id = %v", ground.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed show ground: ground not found: id = %v", ground.Id),
		})
	}

	log.Printf("start show ground: id = %v", ground.Id)

	db.DB.Find(&ground)
	log.Printf("show user: id = %v, Name = %v", ground.Id, ground.Name)

	return c.JSON(ground)
}

/*
	UserUpdate
*/
func GroundUpdate(c *fiber.Ctx) error {
	ground := controllerlogics.GetGroundFromId(c)

	//check account
	err := db.DB.First(&ground).Error
	if err != nil {
		log.Printf("failed update ground: ground not found: id = %v", ground.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed update ground: ground not found: id = %v", ground.Id),
		})
	}

	log.Printf("start update ground: id = %v", ground.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed update ground: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed update ground: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed update ground: you need admin or developer permission",
		})
	}

	err2 := c.BodyParser(&ground)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return err2
	}

	db.DB.Model(&ground).Updates(ground)
	log.Println("success update ground")

	return c.JSON(ground)
}

/*
	UserDelete
*/
func GroundDelete(c *fiber.Ctx) error {
	ground := controllerlogics.GetGroundFromId(c)

	//check account
	err := db.DB.First(&ground).Error
	if err != nil {
		log.Printf("failed delete ground: ground not found: id = %v", ground.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed delete ground: ground not found: id = %v", ground.Id),
		})
	}

	log.Printf("start delete ground: id = %v", ground.Id)

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
		log.Println("failed delete ground: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed delete ground: you need admin or developer permission",
		})
	}

	db.DB.Delete(ground)
	log.Println("success delete ground")

	return c.JSON(fiber.Map{
		"message": "success delete ground",
	})
}
