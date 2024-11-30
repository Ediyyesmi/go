package main

import (
	"fmt"
	"log"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)


type User struct {
	ID       int    `db:"id"`
	Name     string `db:"name"`
	Password string `db:"password"`
}

func main() {
	// PostgreSQL bağlantı ayarları
	dsn := "user=postgres password=12345 dbname=kisiler sslmode=disable"
	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		log.Fatalf("Veritabanına bağlanırken hata oluştu: %v", err)
	}
	defer db.Close()

	
	createTable := `
    (
		id SERIAL PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		password VARCHAR(255) NOT NULL
	);`
	db.Exec(createTable)
	
	insertData := `
	INSERT INTO users (name, password)
	VALUES
		('Seyyide', '111'),
		('Nur', '222');`
	_, err = db.Exec(insertData)
	if err != nil {
		log.Fatalf("Veri eklenirken hata oluştu: %v", err)
	}

	
	var users []User
	err = db.Select(&users, "SELECT * FROM users;")
	if err != nil {
		log.Fatalf("Veri sorgulanırken hata oluştu: %v", err)
	}

	
	fmt.Println("Kayıtlı Kullanıcılar:")
	for _, user := range users {
		fmt.Printf("ID: %d, Name: %s, Password: %s\n", user.ID, user.Name, user.Password)
	}
}
