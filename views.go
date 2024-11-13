package handlers

import (
	"github.com/gofiber/fiber/v2"
)

func RegisterViewHandlers(app *fiber.App) {
	app.Get("/", func(c *fiber.Ctx) error {
		// Render index
		return c.Render("index", fiber.Map{
			"Title": "Home",
		}, "layouts/main")
	})

	app.Get("/layout", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("index", fiber.Map{
			"Title": "Hello, World!",
		}, "layouts/main")
	})

	app.Get("/signup", func(c *fiber.Ctx) error {
		return c.SendFile("views/authentication/register.html")
	})
	app.Get("/crud", func(c *fiber.Ctx) error {
		return c.SendFile("views/crud/crud.html")
	})
	app.Get("/estadisticas", func(c *fiber.Ctx) error {
		return c.SendFile("views/estadisticas/estadisticas.html")
	})
}
