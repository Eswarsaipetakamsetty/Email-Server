package emailhandlers

import (
	"fmt"
	"os"

	"gopkg.in/mail.v2"

	emailmodels "github.com/Eswarsaipetakamsetty/emailserver/email_models"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func SendEmail(ctx *fiber.Ctx) error {
	var payload emailmodels.EmailRequest

	err := godotenv.Load()
	if err != nil {
		return err
	}

	if err := ctx.BodyParser(&payload); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"status": "false",
			"error":  err.Error(),
		})
	}

	m := mail.NewMessage()
	m.SetHeader("From", os.Getenv("SMTP_USER"))
	m.SetHeader("To", payload.To)
	m.SetHeader("Subject", payload.Subject)

	m.SetBody("text/html", fmt.Sprintf(`
	<html>
		<body style="font-family: Arial, sans-serif; text-align: center;">
				<h2>Welcome to Sortminder, %s!</h2>
				<p>Thank you for registering with us.</p>
				<p style="font-size: 16px;">Best Regards,</p>
				<p style="font-size: 16px; font-weight: bold;">Best Wishes, Sortminder</p>
		</body>
	</html>
	`, payload.To))

	port := 587

	d := mail.NewDialer(os.Getenv("SMTP_HOST"), port, os.Getenv("SMTP_USER"), os.Getenv("SMTP_PASSWORD"))

	if err := d.DialAndSend(m); err != nil {
		return err
	}

	return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
		"status":  "true",
		"message": "email sent successfully..",
	})
}
