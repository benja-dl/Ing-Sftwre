package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luccasniccolas/monitorWebApp/data"
	"github.com/luccasniccolas/monitorWebApp/utils"
	"time"
)

func UserSignUp(c *fiber.Ctx, db *sql.DB) error {
	user := new(data.NewUserAccount)
	if err := c.BodyParser(user); err != nil {
		return err
	}

	if utils.CheckEmailExist(user.Email, db) {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{
			"message": "el mail ya esta asociado a una cuenta",
		})
	}

	// Crear la entidad tipo 'user'
	var entityID int
	err := db.QueryRow(`INSERT INTO entities (entity_type) VALUES ('user') RETURNING id`).Scan(&entityID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al crear la entidad de usuario",
		})
	}

	hashPassword, err := utils.HashPassword(user.Password)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error al encriptar la información",
		})
	}

	// Insertar los datos del usuario
	_, err = db.Exec(`INSERT INTO users (entity_id, first_name, last_name, email, password_hash)
		VALUES ($1, $2, $3, $4, $5)`, entityID, user.FirstName, user.LastName, user.Email, hashPassword)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al insertar al usuario en la base de datos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message":    "La cuenta ha sido creada con exito",
		"first name": user.FirstName,
		"last name":  user.LastName,
		"email":      user.Email,
	})
}

func SignIn(c *fiber.Ctx, db *sql.DB) error {
	loginData := new(data.LoginData)
	if err := c.BodyParser(loginData); err != nil {
		return err
	}

	user, err := utils.GetUserByEmail(loginData.Email, db)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "revisar que las credenciales sean correctas",
		})
	}

	if !utils.VerifyPassword(loginData.Password, user.Password) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
			"message": "revisar que las credenciales sean correctas",
		})
	}
	claims := jwt.MapClaims{
		"id":   user.ID,
		"name": user.FirstName,
		"exp":  time.Now().Add(time.Hour * 12).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Sesión iniciada correctamente",
		"token":   t,
	})
}

func Profile(c *fiber.Ctx, db *sql.DB) error {
	userToken := c.Locals("user")
	token := userToken.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	user, err := utils.GetUserByID(int(id), db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "error al obtener datos",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"id":         user.ID,
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
	})

}
