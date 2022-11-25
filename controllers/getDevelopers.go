package controllers

import (
	"GoFiber/config"
	"GoFiber/models"
	"github.com/gofiber/fiber/v2"
)

func GetDevelopers(c *fiber.Ctx) error {
	var developers []models.Developer
	config.DB.Find(&developers)
	return c.JSON(&developers)
}
