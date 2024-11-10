package main

import (
	jwtware "	github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"log"
	"os"
	"v2/internal/database"
	"v2/internal/handlers"
	"v2/internal/middlewares"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")

	// Book
	book := api.Group("/books")
	book.Get("/", handlers.GetBooks)
	book.Get("/:id", handlers.GetBook)
	book.Post("/", middlewares.JWTProtected, handlers.CreateBook)
	book.Patch("/:id", middlewares.JWTProtected, handlers.UpdateBook)
	book.Delete("/:id", middlewares.JWTProtected, handlers.DeleteBook)

	// Auth
	auth := api.Group("/auth")
	auth.Post("/login", handlers.Login)
	auth.Post("/register", handlers.Register)

	log.Fatal(app.Listen("localhost:3000"))
}

func JWTProtected(c *fiber.Ctx) error {
	return jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte(os.Getenv("JWT_SECRET"))},
		ContextKey: "jwt",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Return status 401 and failed authentication error.
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": true,
				"msg":   err.Error(),
			})
		},
	})(c)
}
