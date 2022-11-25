package controllers

import (
	"GoFiber/config"
	"GoFiber/models"
	"github.com/gofiber/fiber/v2"
)

func SaveDevelopers(c *fiber.Ctx) error {
	var developer = new(models.Developer)
	if err := c.BodyParser(developer); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	config.DB.Create(&developer)
	return c.JSON(&developer)
}
