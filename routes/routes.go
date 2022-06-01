package routes

import (
	"fish_go_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Post("/api/users", controllers.UsersCreate)
}
