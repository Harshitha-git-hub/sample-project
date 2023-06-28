package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/harshitha-git-hub/sample-project/database"
)

func main() {
	database.ConnectDb()
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("This is my sample project!")
	})

	app.Listen(":3000")
}
