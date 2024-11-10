package router

import (
	"api-fiber-gorm/handler"
	"api-fiber-gorm/middleware"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"github.com/gofiber/fiber/v2"
)

// SetupRoutes setup router api
func ProductsRoutes(app *fiber.App) {
	api := app.Group("/api", logger.New())

	product := api.Group("/product")
	product.Get("/", handler.GetAllProducts)
	product.Get("/:id", handler.GetProduct)
	product.Post("/", middleware.Protected(), handler.CreateProduct)
	product.Delete("/:id", middleware.Protected(), handler.DeleteProduct)
}
