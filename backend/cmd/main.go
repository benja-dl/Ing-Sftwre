package main

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/luccasniccolas/monitorWebApp/config"
	"github.com/luccasniccolas/monitorWebApp/database"
	"github.com/luccasniccolas/monitorWebApp/handlers"
	"log"
)

func main() {
	// carga de variables de configuracion
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error al cargar las variables de configuración:", err)
	}

	// conexion a la base de datos
	db, err := database.ConnectDatabase(&cfg)
	if err != nil {
		log.Fatal("Error al conectar a la base de datos:", err)
	}
	defer db.Close()

	app := fiber.New()

	// Configura CORS para permitir solicitudes del frontend
	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173", // Cambia esto por la URL del frontend en producción
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	api := app.Group("/api")
	auth := api.Group("/authentication")

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!!!")
	})
	// RUTAS DE AUTENTICACIÓN
	auth.Post("/signup", func(c *fiber.Ctx) error { return handlers.UserSignUp(c, db) })
	auth.Post("/signin", func(c *fiber.Ctx) error { return handlers.SignIn(c, db) })

	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
		ContextKey: "user",
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			log.Println("JWT error:", err) // Esto te ayuda a depurar cualquier problema en el middleware
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"error": "Token missing or malformed",
			})
		},
	}))
	auth.Get("/profile", func(c *fiber.Ctx) error { return handlers.Profile(c, db) })

	err = app.Listen(":8080")
	if err != nil {
		return
	}
}
