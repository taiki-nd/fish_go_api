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
	Index Fish
*/
func FishesIndex(c *fiber.Ctx) error {
	log.Println("get all fishes")

	var fishes []models.Fish
	db.DB.Find(&fishes)

	return c.JSON(fiber.Map{
		"data": fishes,
	})
}

/*
	Create Fish
*/
func FishesCreate(c *fiber.Ctx) error {

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed create fish: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed create fish: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed create fish: you need admin or developer permission",
		})
	}

	log.Println("start to create fish")

	var fish models.Fish

	err := c.BodyParser(&fish)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return err
	}
	db.DB.Create(&fish)
	log.Printf("finish create fish: %v", fish.Fish)

	return c.JSON(fish)
}

/*
	Show Fish
*/
func FishShow(c *fiber.Ctx) error {
	fish := controllerlogics.GetFishFromId(c)

	//check fish
	err := db.DB.First(&fish).Error
	if err != nil {
		log.Printf("failed show fish: fish not found: id = %v", fish.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed show fish: fish not found: id = %v", fish.Id),
		})
	}

	log.Printf("start show fish: id = %v", fish.Id)

	db.DB.Find(&fish)
	log.Printf("show user: id = %v, fish = %v", fish.Id, fish.Fish)

	return c.JSON(fish)
}

/*
	Update Fish
*/
func FishUpdate(c *fiber.Ctx) error {
	fish := controllerlogics.GetFishFromId(c)

	//check account
	err := db.DB.First(&fish).Error
	if err != nil {
		log.Printf("failed update fish: fish not found: id = %v", fish.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed update fish: fish not found: id = %v", fish.Id),
		})
	}

	log.Printf("start update fish: id = %v", fish.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed update fish: please login")
		return c.JSON(fiber.Map{
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed update fish: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed update fish: you need admin or developer permission",
		})
	}

	err2 := c.BodyParser(&fish)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return err2
	}

	db.DB.Model(&fish).Updates(fish)
	log.Println("success update fish")

	return c.JSON(fish)
}

/*
	Delete Fish
*/
func FishDelete(c *fiber.Ctx) error {
	fish := controllerlogics.GetFishFromId(c)

	//check account
	err := db.DB.First(&fish).Error
	if err != nil {
		log.Printf("failed delete fish: fish not found: id = %v", fish.Id)
		return c.JSON(fiber.Map{
			"message": fmt.Sprintf("failed delete fish: fish not found: id = %v", fish.Id),
		})
	}

	log.Printf("start delete fish: id = %v", fish.Id)

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
		log.Println("failed delete fish: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"message": "failed delete fish: you need admin or developer permission",
		})
	}

	db.DB.Delete(fish)
	log.Println("success delete fish")

	return c.JSON(fiber.Map{
		"message": "success delete fish",
	})
}
