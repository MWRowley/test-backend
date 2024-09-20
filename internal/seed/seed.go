package seed

import (
	"log"
	"test-backend/internal/models"
	"test-backend/internal/repositories"
)

func SeedUsers(userRepo *repositories.UserRepository) {
	var count int
	err := userRepo.DB.QueryRow("SELECT COUNT(*) FROM users").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to get user count: %v", err)
	}

	if count > 0 {
		log.Println("Users already exists, skipping seed")
		return
	}

	users := []models.User{
		{Name: "Matthew", Email: "mattrolwey10@gmail.com", Password: "qazwsx123"},
		{Name: "Katie", Email: "katiemorales10@gmail.com", Password: "123456"},
		{Name: "Nick", Email: "Nick@gmail.com", Password: "popeiscools4"},
		{Name: "Brad-Mike", Email: "bradMike@gmail.com", Password: "5dollar"},
		{Name: "Michael", Email: "michael@gmail.com", Password: "bichael"},
		{Name: "Noah", Email: "noah@gmail.com", Password: "gotthatdoginme"},
	}

	for _, user := range users {
		err := userRepo.CreateUser(&user)
		if err != nil {
			log.Printf("Failed to seed user %s: %v", user.Name, err)
		}
		log.Printf("Seeded user %s", user.Name)
	}
}

func SeedPosts(postRepo *repositories.PostRepository) {
	var count int
	err := postRepo.DB.QueryRow("SELECT COUNT(*) FROM posts").Scan(&count)
	if err != nil {
		log.Fatalf("Failed to get post count: %v", err)
	}

	if count > 0 {
		log.Println("Posts already exists, skipping seed")
		return
	}

	posts := []models.Post{
		{Title: "First Post", Content: "This is the first post"},
		{Title: "Second Post", Content: "This is the second post"},
		{Title: "Third Post", Content: "This is the third post"},
		{Title: "Fourth Post", Content: "This is the fourth post"},
		{Title: "Fifth Post", Content: "This is the fifth post"},
		{Title: "Sixth Post", Content: "This is the sixth post"},
	}

	for _, post := range posts {
		err := postRepo.CreatePost(&post)
		if err != nil {
			log.Printf("Failed to seed post %s: %v", post.Title, err)
		}
		log.Printf("Seeded post %s", post.Title)
	}
}
