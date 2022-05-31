package main

import (
	"fish_go_api/config"
	"fish_go_api/routes"
	"fish_go_api/utils"

	"github.com/gofiber/fiber/v2"

	"log"
)

func main() {
	utils.Logging(config.Config.Logfile)
	app := fiber.New()

	routes.Routes(app)

	log.Println("starting server at port:8000")
	app.Listen(":8000")
}
