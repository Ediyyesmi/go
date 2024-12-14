package repositories

import (
	"database/sql"
	"project/models"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (repo *UserRepository) CreateUser(user models.User) error {
	query := "INSERT INTO users (name, password) VALUES ($1, $2)"
	_, err := repo.DB.Exec(query, user.Name, user.Password)
	return err
}

func (repo *UserRepository) GetUserByName(name string) (*models.User, error) {
	query := "SELECT id, name, password FROM users WHERE name = $1"
	row := repo.DB.QueryRow(query, name)
	var user models.User
	err := row.Scan(&user.ID, &user.Name, &user.Password)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
