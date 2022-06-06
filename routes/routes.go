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

	app.Get("api/howtos", controllers.HowtosIndex)
	app.Post("/api/howtos", controllers.HowtosCreate)
	app.Get("/api/howtos/:id", controllers.HowtoShow)
	app.Put("/api/howtos/:id", controllers.HowtoUpdate)
	app.Delete("/api/howtos/:id", controllers.HowtoDelete)

	app.Get("api/fishes", controllers.FishesIndex)
	app.Post("/api/fishes", controllers.FishesCreate)
	app.Get("/api/fishes/:id", controllers.FishShow)
	app.Put("/api/fishes/:id", controllers.FishUpdate)
	app.Delete("/api/fishes/:id", controllers.FishDelete)
}
