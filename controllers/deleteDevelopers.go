package controllers

import (
	"GoFiber/config"
	"GoFiber/models"
	"github.com/gofiber/fiber/v2"
)

func DeleteDeveloper(c *fiber.Ctx) error {
	id := c.Params("id")
	var developer models.Developer
	config.DB.First(&developer, id)
	if developer.Email == "" {
		return c.Status(500).SendString("Developer Not Available")
	}

	config.DB.Delete(&developer)
	return c.SendString("Developer Is Deleted!!!")
}
