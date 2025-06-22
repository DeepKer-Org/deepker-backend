package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

// PaginationParams holds pagination parameters
type PaginationParams struct {
	Page  int `json:"page"`
	Limit int `json:"limit"`
}

// GetPaginationParams extracts and validates pagination parameters from query string
func GetPaginationParams(c *gin.Context) PaginationParams {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}

	limit, err := strconv.Atoi(c.DefaultQuery("limit", "10"))
	if err != nil || limit < 1 {
		limit = 10
	}

	// Limit maximum page size to prevent abuse
	if limit > 100 {
		limit = 100
	}

	return PaginationParams{
		Page:  page,
		Limit: limit,
	}
}

// GetOffset calculates the offset for database queries
func (p PaginationParams) GetOffset() int {
	return (p.Page - 1) * p.Limit
}

// SortParams holds sorting parameters
type SortParams struct {
	SortBy    string `json:"sort_by"`
	SortOrder string `json:"sort_order"`
}

// GetSortParams extracts and validates sort parameters from query string
func GetSortParams(c *gin.Context, defaultSortBy string) SortParams {
	sortBy := c.DefaultQuery("sort_by", defaultSortBy)
	sortOrder := c.DefaultQuery("sort_order", "asc")

	// Validate sort order
	if sortOrder != "asc" && sortOrder != "desc" {
		sortOrder = "asc"
	}

	return SortParams{
		SortBy:    sortBy,
		SortOrder: sortOrder,
	}
}

// BuildOrderClause builds a GORM order clause from sort parameters
func (s SortParams) BuildOrderClause() string {
	if s.SortBy == "" {
		return ""
	}
	return s.SortBy + " " + s.SortOrder
}

// FilterParams holds common filter parameters
type FilterParams struct {
	Search   string            `json:"search"`
	Filters  map[string]string `json:"filters"`
	DateFrom string            `json:"date_from"`
	DateTo   string            `json:"date_to"`
}

// GetFilterParams extracts filter parameters from query string
func GetFilterParams(c *gin.Context) FilterParams {
	filters := make(map[string]string)
	
	// Extract all query parameters as potential filters
	for key, values := range c.Request.URL.Query() {
		if len(values) > 0 && key != "page" && key != "limit" && key != "sort_by" && key != "sort_order" {
			switch key {
			case "search", "date_from", "date_to":
				// These are handled separately
				continue
			default:
				filters[key] = values[0]
			}
		}
	}

	return FilterParams{
		Search:   c.Query("search"),
		Filters:  filters,
		DateFrom: c.Query("date_from"),
		DateTo:   c.Query("date_to"),
	}
}