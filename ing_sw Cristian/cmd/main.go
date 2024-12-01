package main

import (
	"app/config"
	"app/handlers"
	"app/pkg"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {

	// se cargan las variables de entorno, para la conexion a la base de datos
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	// Conectar a la base de datos
	db, err := pkg.ConnectDatabase(cfg)
	if err != nil {
		log.Fatalf("Error connecting to database: %v", err)
	}
	if db != nil {
		log.Println("Conexion a la base de datos exitosa")
	}

	engine := html.New("./views", ".html")

	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/static", "./static")

	pkg.Migrate(db)

	handlers.RegisterViewHandlers(app)
	handlers.RegisterAccountHandlers(app, db)
	handlers.RegisterAuthenticationHandlers(app, db)
	handlers.RegisterDomainHandlers(app, db)

	app.Listen(":" + os.Getenv("SERVER_PORT"))
	log.Println("Server listening on port " + os.Getenv("SERVER_PORT"))
}
