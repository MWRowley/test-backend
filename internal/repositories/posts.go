package repositories

import (
	"database/sql"
	"test-backend/internal/models"
	"time"
)

type Post interface {
	GetPosts() ([]models.Post, error)
	GetPostById(id int) (*models.Post, error)
	CreatePost(post *models.Post) error
	UpdatePost(post *models.Post) error
	DeletePost(id uint) error
}

type PostRepository struct {
	DB *sql.DB
}

func NewPostRepository(DB *sql.DB) *PostRepository {
	return &PostRepository{
		DB: DB,
	}
}

func (r *PostRepository) GetPosts() ([]models.Post, error) {
	rows, err := r.DB.Query("SELECT * FROM posts")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var posts []models.Post
	for rows.Next() {
		var post models.Post
		if err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt); err != nil {
			return nil, err
		}
		posts = append(posts, post)
	}
	return posts, nil
}

func (r *PostRepository) GetPostById(id int) (*models.Post, error) {
	var post models.Post
	if err := r.DB.QueryRow("SELECT * FROM posts WHERE id = $1", id).Scan(&post.Id, &post.Title, &post.Content, &post.CreatedAt); err != nil {
		return nil, err
	}
	return &post, nil
}

func (r *PostRepository) CreatePost(post *models.Post) error {
	post.CreatedAt = time.Now()

	_, err := r.DB.Exec("INSERT INTO posts (title, content, created_at) VALUES ($1, $2, $3)", post.Title, post.Content, post.CreatedAt)
	if err != nil {
		return err
	}
	return err
}

func (r *PostRepository) UpdatePost(post *models.Post) error {
	_, err := r.DB.Exec("UPDATE posts SET title = ?, content = ? WHERE id = ?", post.Title, post.Content)
	if err != nil {
		return err
	}
	return err
}

func (r *PostRepository) DeletePost(id uint) error {
	_, err := r.DB.Exec("DELETE FROM posts WHERE id = ?", id)
	if err != nil {
		return err
	}
	return err
}
