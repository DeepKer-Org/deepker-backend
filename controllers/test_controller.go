package controllers

import (
	"biometric-data-backend/models"
	"biometric-data-backend/services"
	"github.com/gin-gonic/gin"
)

type TestController struct {
	controller *Controller[*models.Test]
}

func NewTestController(service services.Service[*models.Test]) *TestController {
	return &TestController{
		controller: NewController(service),
	}
}

// @Summary Get all tests
// @Description Get all tests
// @Tags tests
// @Accept  json
// @Produce  json
// @Success 200 {array} models.Test
// @Router /tests [get]
func (tc *TestController) GetAllTests(c *gin.Context) {
	tc.controller.GetAll(c)
}

// @Summary Get test by ID
// @Description Get a single test by ID
// @Tags tests
// @Accept  json
// @Produce  json
// @Param   id path int true "Test ID"
// @Success 200 {object} models.Test
// @Router /tests/{id} [get]
func (tc *TestController) GetTestByID(c *gin.Context) {
	tc.controller.GetByID(c)
}

// @Summary Create a new test
// @Description Create a new test
// @Tags tests
// @Accept  json
// @Produce  json
// @Param   test body models.Test true "Test"
// @Success 201 {object} models.Test
// @Router /tests [post]
func (tc *TestController) CreateTest(c *gin.Context) {
	tc.controller.Create(c)
}

// @Summary Update a test
// @Description Update an existing test
// @Tags tests
// @Accept  json
// @Produce  json
// @Param   id path int true "Test ID"
// @Param   test body models.Test true "Test"
// @Success 200 {object} models.Test
// @Router /tests/{id} [put]
func (tc *TestController) UpdateTest(c *gin.Context) {
	tc.controller.Update(c)
}

// @Summary Delete a test
// @Description Delete a test by ID
// @Tags tests
// @Accept  json
// @Produce  json
// @Param   id path int true "Test ID"
// @Success 204
// @Router /tests/{id} [delete]
func (tc *TestController) DeleteTest(c *gin.Context) {
	tc.controller.Delete(c)
}
