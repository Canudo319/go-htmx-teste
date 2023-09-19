package main

import (
	"log"

	"github.com/Canudo319/go-htmx-teste/pkg/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/redirect"
)

func helloWorld(c *fiber.Ctx) error {
	return c.Render("index.html", fiber.Map{
		"PageTitle": "Teste Go HTMX",
		"Itens": []string{
			"Item 1",
			"Item 2",
			"Item 3",
			"Item 4",
			"Item 5",
			"Item 6",
		},
	})
}

func main() {
	app := fiber.New()

	app.Use(redirect.New(redirect.Config{
		Rules: map[string]string{
			"/": "/home",
		},
		StatusCode: 301,
	}))
	app.Get("/home", helloWorld)
	routes.NotFound(app)

	log.Fatal(app.Listen("localhost:8080"))
}
