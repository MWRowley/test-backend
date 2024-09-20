package repositories

import (
	"database/sql"
	"test-backend/internal/models"
)

type Post interface {
	GetPosts() ([]models.Post, error)
	GetPostById(id int) (*models.Post, error)
	UpdatePost(post *models.Post) error
	CreatePost(post *models.Post) error
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
