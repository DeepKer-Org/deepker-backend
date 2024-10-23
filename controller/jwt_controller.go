package controller

import (
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type AuthController struct{}

// NewAuthController creates a new AuthController
func NewAuthController() *AuthController {
	return &AuthController{}
}

// GenerateTokenEndpoint generates a JWT token with the given username and roles
func (ac *AuthController) GenerateTokenEndpoint(c *gin.Context) {
	username := c.PostForm("username")
	roles := c.PostFormArray("roles")

	if username == "" || len(roles) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Email and roles are required"})
		return
	}

	token, err := service.GenerateToken(username, roles)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to generate token"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"token": token})
}
