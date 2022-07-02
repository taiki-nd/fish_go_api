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
	Index ground
*/
func GroundsIndex(c *fiber.Ctx) error {
	log.Println("get all grounds")

	var grounds []models.Ground
	db.DB.Preload("Styles").Preload("Howtos").Preload("Fishes").Find(&grounds)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   grounds,
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
			"status":  false,
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed create ground: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "failed create ground: you need admin or developer permission",
		})
	}

	log.Println("start to create ground")

	var groundAssoci models.GroundAssociation

	err := c.BodyParser(&groundAssoci)
	if err != nil {
		log.Printf("POST method error: %v", err)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("POST method error: %v", err),
		})
	}

	styles := controllerlogics.GetStyles(groundAssoci.Styles)
	howtos := controllerlogics.GetHowtos(groundAssoci.Howtos)
	fishes := controllerlogics.GetFishes(groundAssoci.Fishes)

	ground := models.Ground{
		Name:    groundAssoci.Name,
		Address: groundAssoci.Address,
		Tell:    groundAssoci.Tell,
		Email:   groundAssoci.Email,
		Break:   groundAssoci.Break,
		Styles:  styles,
		Howtos:  howtos,
		Fishes:  fishes,
		Price:   groundAssoci.Price,
		Url:     groundAssoci.Url,
		Feature: groundAssoci.Feature,
		Rule:    groundAssoci.Rule,
		Other:   groundAssoci.Other,
	}

	db.DB.Create(&ground)
	log.Printf("finish create ground: %v", ground.Name)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   ground,
	})
}

/*
	Show Ground
*/
func GroundShow(c *fiber.Ctx) error {
	ground := controllerlogics.GetGroundFromId(c)

	//check ground
	err := db.DB.First(&ground).Error
	if err != nil {
		log.Printf("failed show ground: ground not found: id = %v", ground.Id)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("failed show ground: ground not found: id = %v", ground.Id),
		})
	}

	log.Printf("start show ground: id = %v", ground.Id)

	db.DB.Preload("Styles").Preload("Howtos").Preload("Fishes").Preload("GroundComments").Find(&ground)
	log.Printf("show ground: id = %v, Name = %v", ground.Id, ground.Name)

	return c.JSON(fiber.Map{
		"status": true,
		"data":   ground,
	})
}

/*
	Update ground
*/
func GroundUpdate(c *fiber.Ctx) error {
	ground := controllerlogics.GetGroundFromId(c)

	//check record
	err := db.DB.First(&ground).Error
	if err != nil {
		log.Printf("failed update ground: ground not found: id = %v", ground.Id)
		return c.JSON(fiber.Map{
			"status":  false,
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
			"status":  false,
			"message": "please login",
		})
	}

	//check permission
	var loginUser models.User
	db.DB.Where("id =?", issuer).First(&loginUser)
	if loginUser.PermissionType != "admin" && loginUser.PermissionType != "developer" {
		log.Println("failed update ground: you need admin or developer permission")
		return c.JSON(fiber.Map{
			"status":  false,
			"message": "failed update ground: you need admin or developer permission",
		})
	}

	var groundAssoci models.GroundAssociation

	err2 := c.BodyParser(&groundAssoci)
	if err2 != nil {
		log.Printf("put method error: %v", err2)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("put method error: %v", err2),
		})
	}

	db.DB.Table("ground_styles").Where("ground_id = ?", ground.Id).Delete("")
	db.DB.Table("ground_howtos").Where("ground_id = ?", ground.Id).Delete("")
	db.DB.Table("ground_fishes").Where("ground_id = ?", ground.Id).Delete("")

	styles := controllerlogics.GetStyles(groundAssoci.Styles)
	howtos := controllerlogics.GetHowtos(groundAssoci.Howtos)
	fishes := controllerlogics.GetFishes(groundAssoci.Fishes)

	groundForUpdate := models.Ground{
		Id:      ground.Id,
		Name:    groundAssoci.Name,
		Address: groundAssoci.Address,
		Tell:    groundAssoci.Tell,
		Email:   groundAssoci.Email,
		Break:   groundAssoci.Break,
		Styles:  styles,
		Howtos:  howtos,
		Fishes:  fishes,
		Price:   groundAssoci.Price,
		Url:     groundAssoci.Url,
		Feature: groundAssoci.Feature,
		Rule:    groundAssoci.Rule,
		Other:   groundAssoci.Other,
	}

	db.DB.Model(&groundForUpdate).Updates(groundForUpdate)
	log.Println("success update ground")

	return c.JSON(fiber.Map{
		"status": true,
		"data":   groundForUpdate,
	})
}

/*
	Delete ground
*/
func GroundDelete(c *fiber.Ctx) error {
	ground := controllerlogics.GetGroundFromId(c)

	//check record
	err := db.DB.First(&ground).Error
	if err != nil {
		log.Printf("failed delete ground: ground not found: id = %v", ground.Id)
		return c.JSON(fiber.Map{
			"status":  false,
			"message": fmt.Sprintf("failed delete ground: ground not found: id = %v", ground.Id),
		})
	}

	log.Printf("start delete ground: id = %v", ground.Id)

	// check login or not
	cookie := c.Cookies("jwt")
	issuer, _ := controllerlogics.ParseJwt(cookie)
	if issuer == "" {
		log.Println("failed delete ground: please login")
		return c.JSON(fiber.Map{
			"status":  false,
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

	// delete asociations (transaction)
	err = db.DB.Transaction(func(tx *gorm.DB) error {
		err := tx.Table("ground_styles").Where("ground_id = ?", ground.Id).Delete("").Error
		if err != nil {
			return err
		}

		err = tx.Table("ground_howtos").Where("ground_id = ?", ground.Id).Delete("").Error
		if err != nil {
			return err
		}

		err = tx.Table("ground_fishes").Where("ground_id = ?", ground.Id).Delete("").Error
		if err != nil {
			return err
		}

		var groundComment_id []int64
		tx.Table("ground_comments").Where("ground_id = ?", ground.Id).Pluck("id", &groundComment_id)
		if len(groundComment_id) != 0 {
			err = tx.Table("comment_replies").Where("ground_comment_id IN (?)", groundComment_id).Delete("").Error
			if err != nil {
				return err
			}
		}

		err = tx.Table("ground_comments").Where("ground_id = ?", ground.Id).Delete("").Error
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

	if ground.Filename != "" {
		err := ImageDelete(ground.Filename)
		if err != "" {
			log.Println(err)
			c.JSON(fiber.Map{
				"status":  false,
				"message": err,
			})
		}
	}

	db.DB.Delete(ground)
	log.Println("success delete ground")

	return c.JSON(fiber.Map{
		"status":  true,
		"message": "success delete ground",
	})
}
