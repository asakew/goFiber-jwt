package main

import (
	"api-fiber-gorm/database"
	"api-fiber-gorm/router"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/gofiber/template/html/v2"
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env file
	if err := godotenv.Load(".env"); err != nil {
		panic("Error loading .env file")
	}

	engine := html.New("./web/templates", ".html") // Load templates

	app := fiber.New(fiber.Config{
		Views: engine, // Load templates
	})

	// Serve static files
	app.Static("/assets", "./web/assets")

	// Middleware
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Connect to database
	database.ConnectBookDB()
	database.ConnectPostDB()

	// Routes
	router.HTMLRoutes(app)
	//router.AuthRoutes(app)
	router.ProductsRoutes(app)

	log.Fatal(app.Listen(":3004"))
}
