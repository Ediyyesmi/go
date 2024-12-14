package handlers

import (
	"database/sql"
	"project/repositories"
	"project/services"

	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(app *fiber.App, db *sql.DB) {
	// UserRepository oluştur
	userRepo := repositories.NewUserRepository(db)

	// UserService oluştur
	userService := services.NewUserService(userRepo)

	app.Post("/register", func(c *fiber.Ctx) error {
		var input struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Geçersiz giriş verisi")
		}

		if err := userService.RegisterUser(input.Name, input.Password); err != nil {
			return c.Status(fiber.StatusInternalServerError).SendString("Kayıt hatası")
		}

		return c.Status(fiber.StatusCreated).SendString("Kullanıcı başarıyla kaydedildi")
	})

	app.Post("/login", func(c *fiber.Ctx) error {
		var input struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Geçersiz giriş verisi")
		}

		success, err := userService.AuthenticateUser(input.Name, input.Password)
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).SendString("Kimlik doğrulama hatası")
		}

		if !success {
			return c.Status(fiber.StatusUnauthorized).SendString("Geçersiz kullanıcı adı veya şifre")
		}

		return c.SendString("Başarıyla giriş yapıldı")
	})
}
