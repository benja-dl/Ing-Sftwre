package handlers

import (
	"database/sql"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
	"github.com/luccasniccolas/monitorWebApp/data"
	"github.com/luccasniccolas/monitorWebApp/utils"
	"log"
)

func AddDomain(c *fiber.Ctx, db *sql.DB) error {
	userToken := c.Locals("user")
	token := userToken.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	domain := new(data.NewDomain)
	if err := c.BodyParser(domain); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "campos incorrectos",
		})
	}

	_, err := db.Exec(`INSERT INTO websites (owner_entity_id, name, url)
						VALUES ($1, $2, $3)`, int(id), domain.Name, domain.URL)

	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "no se ha podido agregar el dominio",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "dominio agregado correctamente",
	})

}

func DeleteDomain(c *fiber.Ctx, db *sql.DB) error {
	userToken := c.Locals("user")
	token := userToken.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)
	log.Printf("usuario con id %v eliminando dominio ...", int(id))

	domain := new(data.DomainID)
	err := c.BodyParser(domain)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "campos incorrectos",
		})
	}

	err = utils.DeleteDomain(domain.ID, db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al eliminar el dominio",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "dominio eliminado",
	})

}

func UpdateDomainName(c *fiber.Ctx, db *sql.DB) error {
	domain := new(data.DomainName)
	err := c.BodyParser(domain)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "campos incorrectos",
		})
	}

	err = utils.ChangeDomainName(domain.ID, domain.Name, db)
	if err != nil {
		c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al cambiar el nombre",
		})
	}
	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "dominio editado",
	})
}

func GetAllDomains(c *fiber.Ctx, db *sql.DB) error {
	userToken := c.Locals("user")
	token := userToken.(*jwt.Token)
	claims := token.Claims.(jwt.MapClaims)
	id := claims["id"].(float64)

	domains, err := utils.GetDomains(int(id), db)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Error al obtener dominios",
		})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"domains": domains,
	})
}
