package routes

import (
	"GoFiber/controllers"
	"github.com/gofiber/fiber/v2"
)

func Routes(app *fiber.App) {
	app.Get("/developers", controllers.GetDevelopers)
	app.Get("/developer/:id", controllers.GetDeveloper)
	app.Post("/developers", controllers.SaveDevelopers)
	app.Delete("/developer/:id", controllers.DeleteDeveloper)
	app.Put("/developer/:id", controllers.UpdateDeveloper)

	app.Listen(":7000")

}
