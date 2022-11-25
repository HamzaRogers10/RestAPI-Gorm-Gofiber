package controllers

import (
	"GoFiber/config"
	"GoFiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetDeveloper(c *fiber.Ctx) error {
	id := c.Params("id")
	var developer models.Developer
	config.DB.Find(&developer, id)
	return c.JSON(&developer)
}
