package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sdGuille/trivia-go/database"
)

func main() {
	database.ConnectDb()

	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello Mom i'm Using docker and go, how cool is that!")
	})

	app.Listen(":3000")
}
