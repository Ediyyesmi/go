package main

import (
	"database/sql"
	"log"

	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// Kullanıcı yapısını tanımlıyoruz
type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

func main() {
	// PostgreSQL bağlantısını oluştur
	dsn := "user=postgres password=12345 dbname=kisiler sslmode=disable"
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatalf("Veritabanına bağlanılamadı: %v", err)
	}
	defer db.Close()

	// Fiber uygulamasını oluştur
	app := fiber.New()

	// Anasayfa endpoint'i
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Giriş yapmak için /login endpoint'ine POST isteği gönderin.")
	})

	// Login endpoint'i
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

		// Sonuçları depolayacağımız bir dilim (slice)
		var results []fiber.Map

		// Her kullanıcı için giriş kontrolü yapıyoruz
		for _, input := range inputs {
			var user User
			query := "SELECT id, name, password FROM users WHERE name = $1"
			err := db.QueryRow(query, input.Name).Scan(&user.ID, &user.Name, &user.Password)

			if err != nil {
				// Kullanıcı bulunamazsa
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
				// Kullanıcı bulundu ve şifre doğruysa
				if input.Password == user.Password {
					results = append(results, fiber.Map{
						"name":  input.Name,
						"login": "Başarılı",
					})
				} else {
					// Şifre yanlışsa
					results = append(results, fiber.Map{
						"name":  input.Name,
						"login": "Hatalı şifre",
					})
				}
			}
		}

		// Tüm sonuçları döndür
		return c.JSON(results)
	})

	log.Fatal(app.Listen(":8082"))
}
