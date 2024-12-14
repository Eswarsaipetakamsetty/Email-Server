package main

import (
	"log"

	emailhandlers "github.com/Eswarsaipetakamsetty/emailserver/email_handlers"
	"github.com/Eswarsaipetakamsetty/emailserver/infrastructure"
	"github.com/gofiber/fiber/v2"
)

func main() {

	app := fiber.New()

	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("Build by sortminder")
	})

	infrastructure.InitializePostgresSQL()

	emailhandlers.EmailRoutes(app)

	app.Listen(":8080")
	log.Printf("Server listening on port:8080")
}
