package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
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

	app.Get("/", helloWorld)

	log.Fatal(app.Listen("localhost:8080"))
}
