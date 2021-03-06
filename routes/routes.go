package routes

import (
	"fish_go_api/controllers"

	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("api/users", controllers.UsersIndex)
	app.Post("/api/users", controllers.UsersCreate)
	app.Get("/api/users/:id", controllers.UserShow)
	app.Put("/api/users/:id", controllers.UserUpdate)
	app.Delete("/api/users/:id", controllers.UserDelete)
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

	app.Get("api/ground_comments", controllers.GroundCommentsIndex)
	app.Post("/api/ground_comments", controllers.GroundCommentsCreate)
	app.Get("/api/ground_comments/:id", controllers.GroundCommentShow)
	app.Put("/api/ground_comments/:id", controllers.GroundCommentUpdate)
	app.Delete("/api/ground_comments/:id", controllers.GroundCommentDelete)

	app.Get("api/comment_replies", controllers.CommentRepliesIndex)
	app.Post("/api/comment_replies", controllers.CommentRepliesCreate)
	app.Get("/api/comment_replies/:id", controllers.CommentReplyShow)
	app.Put("/api/comment_replies/:id", controllers.CommentReplyUpdate)
	app.Delete("/api/comment_replies/:id", controllers.CommentReplyDelete)

	app.Get("api/posts", controllers.PostsIndex)
	app.Post("/api/posts", controllers.PostsCreate)
	app.Get("/api/posts/:id", controllers.PostShow)
	app.Put("/api/posts/:id", controllers.PostUpdate)
	app.Delete("/api/posts/:id", controllers.PostDelete)

	app.Get("api/post_comments", controllers.PostCommentsIndex)
	app.Post("/api/post_comments", controllers.PostCommentsCreate)
	app.Get("/api/post_comments/:id", controllers.PostCommentShow)
	app.Put("/api/post_comments/:id", controllers.PostCommentUpdate)
	app.Delete("/api/post_comments/:id", controllers.PostCommentDelete)

	app.Post("api/image", controllers.ImageUpload)
}
