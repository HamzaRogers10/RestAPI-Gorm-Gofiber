package controllers

import (
	"GoFiber/config"
	"GoFiber/models"
	"github.com/gofiber/fiber/v2"
)

func UpdateDeveloper(c *fiber.Ctx) error {
	id := c.Params("id")
	var developer = models.Developer{}
	config.DB.First(&developer, id)
	if developer.Email == "" {
		return c.Status(500).SendString("User not available")
	}
	if err := c.BodyParser(&developer); err != nil {
		return c.Status(500).SendString(err.Error())
	}
	config.DB.Save(&developer)
	return c.JSON(&developer)
}
