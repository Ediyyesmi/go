// şifreleri postgreste hash'leyerek tutar

package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // PostgreSQL 
	"golang.org/x/crypto/bcrypt" //hash
)


type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

func main() {
	
	dsn := "user=postgres password=12345 dbname=kisiler sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}
	defer db.Close()
	
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Giriş yapmak için /login endpoint'ine POST isteği gönderin.")
	})


	app.Post("/register", func(c *fiber.Ctx) error {
		var input struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		// JSON verisini çözümle
		if err := c.BodyParser(&input); err != nil {
			return c.Status(fiber.StatusBadRequest).SendString("Geçersiz giriş verisi")
		}

		//şifreyi hash'le
		hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
		if err != nil {
			log.Println("Şifre hashleme hatası:", err)
			return c.Status(fiber.StatusInternalServerError).SendString("Kullanıcı kaydı sırasında hata oluştu")
		}

		
		queryInsert := "INSERT INTO users (name, password) VALUES ($1, $2)"
		db.Exec(queryInsert, input.Name, string(hashedPassword)) 

		return c.Status(fiber.StatusCreated).SendString("Kullanıcı başarıyla kaydedildi")
	})

	
	app.Post("/login", func(c *fiber.Ctx) error {
		var inputs []struct {
			Name     string `json:"name"`
			Password string `json:"password"`
		}

		// JSON verisini çözümle
		if err := c.BodyParser(&inputs); err != nil {
			log.Println("JSON çözümleme hatası:", err) // Hata mesajını logla
			return c.Status(fiber.StatusBadRequest).SendString("Geçersiz giriş verisi")
		}

		// Sonuçları depola
		var results []fiber.Map

		for _, input := range inputs {
			var user User
			query := "SELECT id, name, password FROM users WHERE name = $1"
			err := db.QueryRow(query, input.Name).Scan(&user.ID, &user.Name, &user.Password)

			if err != nil {
				
				if err == sql.ErrNoRows {
					results = append(results, fiber.Map{
						"name":  input.Name,
						"login": "Kullanıcı bulunamadı",
					})
				} else {
					// Diğer veritabanı hatalarını handle et
					return c.Status(fiber.StatusInternalServerError).SendString("Veritabanı hatası")
				}
			} else {
				
				err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password))
				if err != nil {
					results = append(results, fiber.Map{
						"name":  input.Name,
						"login": "Hatalı şifre",
					})
				} else {
					
					results = append(results, fiber.Map{
						"name":  input.Name,
						"login": "Başarılı",
					})
				}
			}
		}

		return c.JSON(results)
	})

	log.Fatal(app.Listen(":8082"))
}
