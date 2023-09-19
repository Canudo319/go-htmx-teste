package routes

import "github.com/gofiber/fiber/v2"

func NotFound(app *fiber.App) {
	app.Use(
		func(c *fiber.Ctx) error {
			return c.Status(fiber.StatusNotFound).Render("pkg/templates/not_found.html", c.Path())
		},
	)
}
