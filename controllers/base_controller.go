package controllers

import (
	"biometric-data-backend/models"
	"biometric-data-backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type Controller[T models.Identifiable] struct {
	service service.Service[T]
}

func NewController[T models.Identifiable](service service.Service[T]) *Controller[T] {
	return &Controller[T]{service: service}
}

func (ctrl *Controller[T]) GetAll(c *gin.Context) {
	entities, err := ctrl.service.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entities)
}

func (ctrl *Controller[T]) GetByID(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	entity, err := ctrl.service.GetByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, entity)
}

func (ctrl *Controller[T]) Create(c *gin.Context) {
	var entity T
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	createdEntity, err := ctrl.service.Create(&entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, createdEntity)
}

func (ctrl *Controller[T]) Update(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	var entity T
	if err := c.ShouldBindJSON(&entity); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	entity.SetID(uint(id))
	updatedEntity, err := ctrl.service.Update(&entity)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, updatedEntity)
}

func (ctrl *Controller[T]) Delete(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	if err := ctrl.service.Delete(uint(id)); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
