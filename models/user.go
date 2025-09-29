package models

import (
	"time"
)

type User struct {
	ID           string    `json:"id" db:"id"`
	Name         string    `json:"name" db:"name"`
	PasswordHash string    `json:"-" db:"password_hash"`
	ImagePath    string    `json:"imagePath" db:"image_path"`
	Content      string    `json:"content" db:"content"`
	CreatedAt    time.Time `json:"createdAt" db:"created_at"`
}

type UserCreateRequest struct {
	Name      string `json:"name"`
	Password  string `json:"password"`
	ImagePath string `json:"imagePath"`
	Content   string `json:"content"`
}
