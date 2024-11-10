package router

import "github.com/gofiber/fiber/v2"

func HTMLRoutes(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Render("index", nil)
	})

	app.Get("/auth/login", func(c *fiber.Ctx) error {
		return c.Render("login", nil)
	})
}
