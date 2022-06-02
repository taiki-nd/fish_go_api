package routes

import (
	"fish_go_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("api/users", controllers.UsersIndex)
	app.Post("/api/users", controllers.UsersCreate)
	app.Get("/api/user/:id", controllers.UserShow)
	app.Put("/api/user/:id", controllers.UserUpdate)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
}
