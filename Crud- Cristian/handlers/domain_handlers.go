package handlers

import (
    "fmt"
    "log"

    "github.com/gofiber/fiber/v2"
    "gorm.io/gorm"
)

type DomainRequest struct {
    Domain    string `json:"domain"`
    OldDomain string `json:"oldDomain,omitempty"`
    NewDomain string `json:"newDomain,omitempty"`
}

// Registrar las rutas para la gestión de dominios
func RegisterDomainHandlers(app *fiber.App, db *gorm.DB) {
    app.Post("/add", func(c *fiber.Ctx) error {
        var req DomainRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        if err := addDomain(db, req.Domain); err != nil {
            log.Println("Error adding domain:", err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Dominio agregado"})
    })

    app.Post("/edit", func(c *fiber.Ctx) error {
        var req DomainRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        if err := editDomain(db, req.OldDomain, req.NewDomain); err != nil {
            log.Println("Error editing domain:", err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Dominio editado"})
    })

    app.Post("/delete", func(c *fiber.Ctx) error {
        var req DomainRequest
        if err := c.BodyParser(&req); err != nil {
            return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
        }
        if err := deleteDomain(db, req.Domain); err != nil {
            log.Println("Error deleting domain:", err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Dominio eliminado"})
    })

    app.Get("/list", func(c *fiber.Ctx) error {
        domains, err := listDomains(db)
        if err != nil {
            log.Println("Error listing domains:", err)
            return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
        }
        return c.JSON(domains)
    })
}

func addDomain(db *gorm.DB, domain string) error {
    return db.Create(&Domain{Name: domain}).Error
}

func deleteDomain(db *gorm.DB, domain string) error {
    result := db.Where("name = ?", domain).Delete(&Domain{})
    if result.RowsAffected == 0 {
        return fmt.Errorf("no se encontró el dominio %s", domain)
    }
    return result.Error
}

func editDomain(db *gorm.DB, oldDomain, newDomain string) error {
    result := db.Model(&Domain{}).Where("name = ?", oldDomain).Update("name", newDomain)
    if result.RowsAffected == 0 {
        return fmt.Errorf("no se encontró el dominio %s", oldDomain)
    }
    return result.Error
}

func listDomains(db *gorm.DB) ([]string, error) {
    var domains []Domain
    if err := db.Find(&domains).Error; err != nil {
        return nil, err
    }
    var names []string
    for _, domain := range domains {
        names = append(names, domain.Name)
    }
    return names, nil
}

// Definir la estructura de dominio para GORM
type Domain struct {
    ID   uint   `gorm:"primaryKey"`
    Name string `gorm:"unique"`
}
