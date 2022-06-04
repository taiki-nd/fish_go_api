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
	app.Delete("/api/user/:id", controllers.UserDelete)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("api/grounds", controllers.GroundsIndex)
	app.Post("/api/grounds", controllers.GroundsCreate)
	app.Get("/api/grounds/:id", controllers.GroundShow)
	app.Put("/api/grounds/:id", controllers.GroundUpdate)
	app.Delete("/api/grounds/:id", controllers.GroundDelete)

	app.Get("api/styles", controllers.StylesIndex)
	app.Post("/api/styles", controllers.StylesCreate)
	app.Get("/api/styles/:id", controllers.StyleShow)
	app.Put("/api/styles/:id", controllers.StyleUpdate)
	app.Delete("/api/styles/:id", controllers.StyleDelete)
}
