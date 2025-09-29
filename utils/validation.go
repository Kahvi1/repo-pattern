package utils

import (
	"errors"
	"mime/multipart"
	"path/filepath"
	"strings"
)

// ValidateName validates user name
func ValidateName(name string) error {
	name = strings.TrimSpace(name)
	if name == "" {
		return errors.New("name is required")
	}
	if len(name) < 2 {
		return errors.New("name must be at least 2 characters")
	}
	if len(name) > 255 {
		return errors.New("name must not exceed 255 characters")
	}
	return nil
}

// ValidatePassword validates user password
func ValidatePassword(password string) error {
	if password == "" {
		return errors.New("password is required")
	}
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters")
	}
	if len(password) > 72 {
		return errors.New("password must not exceed 72 characters")
	}
	return nil
}

// ValidateContent validates user content
func ValidateContent(content string) error {
	content = strings.TrimSpace(content)
	if content == "" {
		return errors.New("content is required")
	}
	if len(content) > 5000 {
		return errors.New("content must not exceed 5000 characters")
	}
	return nil
}

// ValidateImageFile validates uploaded image file
func ValidateImageFile(fileHeader *multipart.FileHeader) error {
	if fileHeader == nil {
		return errors.New("image file is required")
	}

	// Check file size (max 5MB)
	const maxSize = 5 * 1024 * 1024 // 5MB in bytes
	if fileHeader.Size > maxSize {
		return errors.New("image file must not exceed 5MB")
	}

	// Check file extension
	ext := strings.ToLower(filepath.Ext(fileHeader.Filename))
	allowedExts := map[string]bool{
		".jpg":  true,
		".jpeg": true,
		".png":  true,
		".gif":  true,
	}

	if !allowedExts[ext] {
		return errors.New("only .jpg, .jpeg, .png, and .gif files are allowed")
	}

	return nil
}
