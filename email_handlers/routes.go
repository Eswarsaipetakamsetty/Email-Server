package emailhandlers

import "github.com/gofiber/fiber/v2"

func EmailRoutes(incommingRoutes *fiber.App) {
	incommingRoutes.Post("/send-email", SendEmail)
}
