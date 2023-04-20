package handlers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/sdGuille/trivia-go/database"
	"github.com/sdGuille/trivia-go/models"
)


func Home(c *fiber.Ctx) error {
	return c.SendString("Hello Mom i'm Using docker and go, how cool is that!")
}

func ListFacts(c *fiber.Ctx) error {
	facts := []models.Fact{}
	database.DB.Db.Find(&facts)
	
	return c.Render("index", fiber.Map{
		"Title": "sdGuille Trivia Time",
		"subtitle": "Facts for funtimes	with friends!",
	})
}

func CreateFact(c *fiber.Ctx) error {
	fact := new(models.Fact)
	if err := c.BodyParser(fact); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	database.DB.Db.Create(&fact)

	return c.Status(200).JSON(fact)
}
