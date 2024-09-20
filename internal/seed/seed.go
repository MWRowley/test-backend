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
