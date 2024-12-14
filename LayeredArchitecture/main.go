package main

import (
	"database/sql"
	"log"
	"project/handlers"
	"project/repositories"
	"project/services"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

func main() {
	// Veritabanı bağlantısı
	dsn := "user=postgres password=12345 dbname=kisiler sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	defer db.Close()

	// Katmanları başlat
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)
	userHandler := handlers.NewUserHandler(userService)

	// Fiber uygulamasını başlat
	app := fiber.New()

	app.Post("/register", userHandler.Register)
	app.Post("/login", userHandler.Login)

	log.Fatal(app.Listen(":8082"))
}
