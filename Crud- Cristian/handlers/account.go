package handlers

import (
	"app/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func RegisterAccountHandlers(app *fiber.App, db *gorm.DB) {
	app.Get("account", func(c *fiber.Ctx) error {
		var users []models.User
		db.Find(&users)

		return c.JSON(users)
	})

	app.Get("account/:id", func(c *fiber.Ctx) error {
		var user models.User
		db.First(&user, c.Params("id"))

		if user.ID == 0 {
			return c.JSON(fiber.Map{"message": "User not found"})
		}

		return c.JSON(fiber.Map{
			"ID":         user.ID,
			"first_name": user.FirstName,
			"last_name":  user.LastName,
			"mail":       user.Mail,
		})
	})
}
