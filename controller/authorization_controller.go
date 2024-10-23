package controller

import (
	"biometric-data-backend/models/dto"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
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
