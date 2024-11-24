// Fiber framework'ü kullanarak basit bir login (giriş) API'si 

package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	const userID = "seyyide"
	const userPassword = "7878"

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Giriş yapmak için /login endpoint'ine POST isteği gönderin.")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		var input struct {
			ID       string `json:"ID"`
			Password string `json:"Password"`
		}

		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Geçersiz giriş verisi")
		}

		if input.ID == userID && input.Password == userPassword {
			return c.SendString("Giriş başarılı!")
		}

		return c.Status(fiber.StatusUnauthorized).SendString("Giriş başarısız!")
	})

	log.Fatal(app.Listen(":8081"))
}
