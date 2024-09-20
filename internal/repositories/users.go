package repositories

import (
	"database/sql"
	"test-backend/internal/models"
)

type User interface {
	GetUsers() ([]models.User, error)
	GerUserByName(name string) (*models.User, error)
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

func (r *UserRepository) CreateUser(user *models.User) error {
	_, err := r.DB.Exec("INSERT INTO users (id, name, email, password, created_at) VALUES (?, ?, ?, ?, ?)", user.Id)
	if err != nil {
		return nil
	}
	return err
}

func (r *UserRepository) UpdateUser(user *models.User) error {
	_, err := r.DB.Exec("UPDATE users SET name = ?, email = ?, password = ? WHERE id = ?", user.Name, user.Id)
	if err != nil {
		return nil
	}
	return err
}

func (r *UserRepository) DeleteUser(id uint) error {
	_, err := r.DB.Exec("DELETE FROM users WHERE id = ?", id)
	if err != nil {
		return nil
	}
	return err
}
