package repositories

import (
	"database/sql"
	"test-backend/internal/models"
	"time"
)

type User interface {
	GetUsers() ([]models.User, error)
	GetUserByName(name string) (*models.User, error)
	UpdateUser(user *models.User) error
	CreateUser(user *models.User) error
	DeleteUser(id uint) error
}

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(DB *sql.DB) *UserRepository {
	return &UserRepository{
		DB: DB,
	}
}

func (r *UserRepository) GetUsers() ([]models.User, error) {
	rows, err := r.DB.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func (r *UserRepository) GetUserByName(name string) (*models.User, error) {
	var user models.User
	err := r.DB.QueryRow("SELECT * FROM users WHERE name = $1", name).Scan(&user.Id, &user.Name, &user.Email, &user.Password, &user.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(user *models.User) error {
	user.CreatedAt = time.Now()

	_, err := r.DB.Exec("INSERT INTO users (name, email, password, created_at) VALUES ($1, $2, $3, $4)", user.Name, user.Email, user.Password, user.CreatedAt)
	if err != nil {
		return err
	}
	return err
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	_, err := r.DB.Exec("UPDATE users SET name = $1, email = $2, password = $3 WHERE id = $4", user.Name, user.Email, user.Password, user.Id)
	if err != nil {
		return err
	}
	return err
}

func (r *UserRepository) DeleteUser(id uint) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		return err
	}
	return err
}
