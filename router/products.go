package router

import (
	"api-fiber-gorm/handlers"
	"api-fiber-gorm/middlewares"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func ProductsRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	product := api.Group("/product")
	product.Get("/", handlers.GetAllProducts)
	product.Get("/:id", handlers.GetProduct)
	product.Post("/", middlewares.Protected(), handlers.CreateProduct)
	product.Delete("/:id", middlewares.Protected(), handlers.DeleteProduct)
}
