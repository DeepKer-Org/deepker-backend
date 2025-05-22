package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"log"
	"net/http"
)

// Generic function to retrieve a resource by ID (UUID)
func getByID[T any](c *gin.Context, idParam string, fetchFunc func(uuid.UUID) (*T, error), notFoundMessage string) (*T, error) {
	id := c.Param(idParam)

	// Parse the string to a UUID
	parsedID, err := uuid.Parse(id)
	if err != nil {
		log.Printf("Invalid UUID: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID format"})
		return nil, err
	}

	// Fetch the resource using the provided function
	resource, err := fetchFunc(parsedID)
	if err != nil {
		log.Printf("Error retrieving resource: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve resource"})
		return nil, err
	}

	if resource == nil {
		log.Printf(notFoundMessage, id)
		c.JSON(http.StatusNotFound, gin.H{"error": notFoundMessage})
		return nil, nil
	}

	return resource, nil
}

// Generic function to handle JSON binding
func bindJSON[T any](c *gin.Context, payload *T) bool {
	if err := c.ShouldBindJSON(payload); err != nil {
		log.Printf("Error binding JSON: %v", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return false
	}
	return true
}
