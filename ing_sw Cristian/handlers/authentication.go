package handlers

import (
	"app/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
	"log"
)

func RegisterAuthenticationHandlers(app *fiber.App, db *gorm.DB) {
	app.Post("/authentication/signup", func(c *fiber.Ctx) error {
		// Parsear el cuerpo de la solicitud a la estructura UserRegisterData
		u := new(models.UserRegisterData)

		if err := c.BodyParser(u); err != nil {
			log.Println("Error al parsear el cuerpo de la solicitud:", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": "Datos inválidos",
			})
		}

		// Crear el nuevo usuario en la base de datos
		if err := models.CreateNewUser(u, db); err != nil {
			log.Printf("Error al crear usuario: %v\n", err)
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"message": err.Error(),
			})
		}

		// Redirigir al index si el registro fue exitoso
		log.Println("Usuario creado correctamente")
		return c.Redirect("/")
	})
}
