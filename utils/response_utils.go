package utils

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// ErrorResponse represents a standardized error response
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
	Code    string `json:"code,omitempty"`
}

// SuccessResponse represents a standardized success response
type SuccessResponse struct {
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

// RespondWithError sends a standardized error response
func RespondWithError(c *gin.Context, statusCode int, err string, message ...string) {
	response := ErrorResponse{
		Error: err,
	}
	
	if len(message) > 0 {
		response.Message = message[0]
	}

	// Log the error
	log.Printf("Error [%d]: %s", statusCode, err)
	if len(message) > 0 {
		log.Printf("Message: %s", message[0])
	}

	c.JSON(statusCode, response)
}

// RespondWithSuccess sends a standardized success response
func RespondWithSuccess(c *gin.Context, statusCode int, message string, data ...interface{}) {
	response := SuccessResponse{
		Message: message,
	}
	
	if len(data) > 0 {
		response.Data = data[0]
	}

	c.JSON(statusCode, response)
}

// RespondWithData sends a response with just data (no message wrapper)
func RespondWithData(c *gin.Context, statusCode int, data interface{}) {
	c.JSON(statusCode, data)
}

// ValidateUUID validates and parses UUID from string
func ValidateUUID(c *gin.Context, idStr string) (uuid.UUID, bool) {
	id, err := uuid.Parse(idStr)
	if err != nil {
		RespondWithError(c, http.StatusBadRequest, "Invalid UUID format", "The provided ID is not a valid UUID")
		return uuid.Nil, false
	}
	return id, true
}

// BindJSONAndValidate binds JSON request and validates it
func BindJSONAndValidate[T any](c *gin.Context, obj *T) bool {
	if err := c.ShouldBindJSON(obj); err != nil {
		RespondWithError(c, http.StatusBadRequest, "Invalid request payload", err.Error())
		return false
	}
	return true
}

// Common error responses
func RespondNotFound(c *gin.Context, resource string) {
	RespondWithError(c, http.StatusNotFound, resource+" not found")
}

func RespondInternalError(c *gin.Context, operation string, err error) {
	log.Printf("Internal error during %s: %v", operation, err)
	RespondWithError(c, http.StatusInternalServerError, "Internal server error", "Failed to "+operation)
}

func RespondBadRequest(c *gin.Context, message string) {
	RespondWithError(c, http.StatusBadRequest, "Bad request", message)
}

func RespondUnauthorized(c *gin.Context) {
	RespondWithError(c, http.StatusUnauthorized, "Unauthorized", "Authentication required")
}

func RespondForbidden(c *gin.Context) {
	RespondWithError(c, http.StatusForbidden, "Forbidden", "Insufficient permissions")
}