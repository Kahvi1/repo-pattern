package handlers

import (
	"context"
	"net/http"
	"repo_pattern/models"
	"repo_pattern/repository"
	"repo_pattern/utils"
	"time"
)

// UserHandler handles user-related HTTP requests
type UserHandler struct {
	userRepo repository.UserRepository
}

// NewUserHandler creates a new user handler
func NewUserHandler(userRepo repository.UserRepository) *UserHandler {
	return &UserHandler{
		userRepo: userRepo,
	}
}

// CreateUser handles POST /users - User signup
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	// Only accept POST method
	if r.Method != http.MethodPost {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Parse multipart form (max 10MB in memory)
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Failed to parse form data")
		return
	}

	// Extract form fields
	name := r.FormValue("name")
	password := r.FormValue("password")
	content := r.FormValue("content")

	// Validate name
	if err := utils.ValidateName(name); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate password
	if err := utils.ValidatePassword(password); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Validate content
	if err := utils.ValidateContent(content); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Get uploaded file
	_, fileHeader, err := r.FormFile("image")
	if err != nil {
		utils.RespondError(w, http.StatusBadRequest, "Image file is required")
		return
	}

	// Validate image file
	if err := utils.ValidateImageFile(fileHeader); err != nil {
		utils.RespondError(w, http.StatusBadRequest, err.Error())
		return
	}

	// Save uploaded file
	imagePath, err := utils.SaveUploadedFile(fileHeader)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to save image file")
		return
	}

	// Create user object
	user := &models.User{
		Name:         name,
		PasswordHash: password, // Will be hashed in repository
		ImagePath:    imagePath,
		Content:      content,
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Save to database
	if err := h.userRepo.Create(ctx, user); err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to create user")
		return
	}

	// Return success response
	utils.RespondSuccess(w, http.StatusCreated, user, "User created successfully")
}

// GetUserByID handles GET /users/{id}
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	// Only accept GET method
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Extract ID from URL path
	// Example: /users/550e8400-e29b-41d4-a716-446655440000
	id := r.URL.Path[len("/users/"):]
	if id == "" {
		utils.RespondError(w, http.StatusBadRequest, "User ID is required")
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get user from database
	user, err := h.userRepo.GetByID(ctx, id)
	if err != nil {
		utils.RespondError(w, http.StatusNotFound, "User not found")
		return
	}

	// Return success response
	utils.RespondSuccess(w, http.StatusOK, user, "User retrieved successfully")
}

// GetAllUsers handles GET /users
func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	// Only accept GET method
	if r.Method != http.MethodGet {
		utils.RespondError(w, http.StatusMethodNotAllowed, "Method not allowed")
		return
	}

	// Create context with timeout
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Get all users from database
	users, err := h.userRepo.GetAll(ctx)
	if err != nil {
		utils.RespondError(w, http.StatusInternalServerError, "Failed to retrieve users")
		return
	}

	// Return success response
	utils.RespondSuccess(w, http.StatusOK, users, "Users retrieved successfully")
}
