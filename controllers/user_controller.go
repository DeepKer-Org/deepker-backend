package controllers

import (
	"biometric-data-backend/models"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	controller *Controller[*models.User]
}

func NewUserController(service service.Service[*models.User]) *UserController {
	return &UserController{
		controller: NewController(service),
	}
}

// @Summary Get all users
// @Description Get all users
// @Tags users
// @Accept  json
// @Produce  json
// @Success 200 {array} models.User
// @Router /users [get]
func (uc *UserController) GetAllUsers(c *gin.Context) {
	uc.controller.GetAll(c)
}

// @Summary Get user by ID
// @Description Get a single user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id path int true "User ID"
// @Success 200 {object} models.User
// @Router /users/{id} [get]
func (uc *UserController) GetUserByID(c *gin.Context) {
	uc.controller.GetByID(c)
}

// @Summary Create a new user
// @Description Create a new user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   user body models.User true "User"
// @Success 201 {object} models.User
// @Router /users [post]
func (uc *UserController) CreateUser(c *gin.Context) {
	uc.controller.Create(c)
}

// @Summary Update a user
// @Description Update an existing user
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id path int true "User ID"
// @Param   user body models.User true "User"
// @Success 200 {object} models.User
// @Router /users/{id} [put]
func (uc *UserController) UpdateUser(c *gin.Context) {
	uc.controller.Update(c)
}

// @Summary Delete a user
// @Description Delete a user by ID
// @Tags users
// @Accept  json
// @Produce  json
// @Param   id path int true "User ID"
// @Success 204
// @Router /users/{id} [delete]
func (uc *UserController) DeleteUser(c *gin.Context) {
	uc.controller.Delete(c)
}
