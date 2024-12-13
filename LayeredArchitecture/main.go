package main

import (
	"database/sql"
	"log"
	"project/handlers"

	_ "github.com/lib/pq" 
	"github.com/gofiber/fiber/v2"
)

func main(){
	app:= fiber.New()

	db, err := sql.Open("postgres", "user=postgres password=12345 dbname=kisiler sslmode=disable")
	if err != nil {
		log.Fatal("veritabanına bağlanılamadı: %v", err)
	}
	defer db.Close()

	handlers.SetupRoutes(app,db)

	log.Fatal(app.Listen(":8082"))
}