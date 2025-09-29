package main

import (
	"context"
	"log"
	"repo_pattern/database"
	"repo_pattern/models"
	"repo_pattern/repository"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	pool, err := database.NewPostgresDB(ctx)
	if err != nil {
		log.Fatalf("Failed to connect to database: %v\n", err)
	}
	defer pool.Close()

	userRepo := repository.NewPostgresUserRepository(pool)

	testUser := &models.User{
		Name:         "Test pengguna",
		PasswordHash: "password123",
		ImagePath:    "/uplodas/images/test.jpg",
		Content:      "this is my first post",
	}

	if err := userRepo.Create(ctx, testUser); err != nil {
		log.Fatalf("Failed to create user: %v\n", err)
	}

	log.Printf("✅ User created successfully!")
	log.Printf("   ID: %s", testUser.ID)
	log.Printf("   Name: %s", testUser.Name)
	log.Printf("   Created At: %s", testUser.CreatedAt)

	retrievedUser, err := userRepo.GetByID(ctx, testUser.ID)
	if err != nil {
		log.Fatalf("Failed to retrieve user: %v\n", err)
	}

	log.Printf("✅ User retrieved successfully!")
	log.Printf("   ID: %s", retrievedUser.ID)
	log.Printf("   Name: %s", retrievedUser.Name)
	log.Printf("   Content: %s", retrievedUser.Content)

	println("🚀 Application started successfully!")
}
