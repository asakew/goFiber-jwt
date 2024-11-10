package router

import (
	"api-fiber-gorm/handlers"
	"api-fiber-gorm/middlewares"
	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func AuthRoutes(app *fiber.App) {
	// JWT Middleware for secure routes
	app.Use("/admin", middlewares.JwtMiddleware())

	// Setup routes
	app.Get("/admin/posts", handlers.GetPosts)
	app.Post("/admin/posts", handlers.CreatePost)
	app.Put("/admin/posts/:id", handlers.UpdatePost)
	app.Delete("/admin/posts/:id", handlers.DeletePost)
}
