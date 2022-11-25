package main

import (
	"GoFiber/config"
	"GoFiber/routes"
	"github.com/gofiber/fiber/v2"
)

func hello(c *fiber.Ctx) error {
	return c.SendString("Welcome to Developer API!!")
}

func main() {

	config.Migration()
	app := fiber.New()
	app.Get("/", hello)
	routes.Routes(app)
}
