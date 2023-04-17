package handlers

import "github.com/gofiber/fiber/v2"

func Home(c *fiber.Ctx) error {
	return c.SendString("Hello Mom i'm Using docker and go, how cool is that!")
}
