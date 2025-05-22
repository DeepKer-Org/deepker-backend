package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
	"strconv"
)

type AuthorizationController struct {
	UserService service.AuthorizationService
}

func NewAuthorizationController(userService service.AuthorizationService) *AuthorizationController {
	return &AuthorizationController{
		UserService: userService,
	}
}

// AuthenticateUser handles user login and returns a token if successful
func (uc *AuthorizationController) AuthenticateUser(c *gin.Context) {
	var loginDTO dto.UserLoginDTO
	if err := c.ShouldBindJSON(&loginDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	token, err := uc.UserService.AuthenticateUser(&loginDTO)
	if err != nil {
		log.Printf("Authentication failed: %v", err)
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Authentication failed"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Authentication successful", "token": token})
}

// RegisterUser handles user registration
func (uc *AuthorizationController) RegisterUser(c *gin.Context) {
	var userDTO dto.UserRegisterDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	_, err := uc.UserService.RegisterUser(&userDTO)
	if err != nil {
		log.Printf("Failed to register user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to register user"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User registered successfully"})
}

func (uc *AuthorizationController) GetUserById(c *gin.Context) {
	id := c.Param("id")
	userId, err := uuid.Parse(id)

	user, err := uc.UserService.GetUserById(userId)
	if err != nil {
		log.Printf("Failed to get user: %v", err)
		c.JSON(http.StatusNotFound, gin.H{"error": "User not found"})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (uc *AuthorizationController) GetAllUsers(c *gin.Context) {
	var users []*dto.UserDTO
	var totalCount int
	var err error

	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	users, totalCount, err = uc.UserService.GetAllUsers(page, limit)
	if err != nil {
		log.Printf("Failed to get users: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get users"})
		return
	}

	// Include totalCount in the response for status-specific queries
	c.JSON(http.StatusOK, gin.H{
		"users":      users,
		"totalCount": totalCount,
	})
	return
}

func (uc *AuthorizationController) UpdateUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := uuid.Parse(id)

	var userDTO dto.UserUpdateDTO
	if err := c.ShouldBindJSON(&userDTO); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	err = uc.UserService.UpdateUser(userId, &userDTO)
	if err != nil {
		log.Printf("Failed to update user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User updated successfully"})
}

func (uc *AuthorizationController) DeleteUser(c *gin.Context) {
	id := c.Param("id")
	userId, err := uuid.Parse(id)

	err = uc.UserService.DeleteUser(userId)
	if err != nil {
		log.Printf("Failed to delete user: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete user"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully"})
}
