package controllerlogics

import (
	"fish_go_api/models"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func GetGroundFromId(c *fiber.Ctx) models.Ground {
	id, _ := strconv.Atoi(c.Params("id"))
	ground := models.Ground{
		Id: uint(id),
	}

	return ground
}

func GetStyles(stylesInfo []int) []models.Style {
	styles := make([]models.Style, len(stylesInfo))
	for i, styleId := range stylesInfo {
		styles[i] = models.Style{
			Id: uint(styleId),
		}
	}
	return styles
}

func GetHowtos(howtosInfo []int) []models.Howto {
	howtos := make([]models.Howto, len(howtosInfo))
	for i, howtoId := range howtosInfo {
		howtos[i] = models.Howto{
			Id: uint(howtoId),
		}
	}
	return howtos
}

func GetFishes(fishesInfo []int) []models.Fish {
	fishes := make([]models.Fish, len(fishesInfo))
	for i, fishId := range fishesInfo {
		fishes[i] = models.Fish{
			Id: uint(fishId),
		}
	}
	return fishes
}
