package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

func AddUserDomain(c *fiber.Ctx, db *sql.DB) error {
	userToken := c.Locals("user")
	token := userToken.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
}
