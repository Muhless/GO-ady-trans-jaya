package controllers

import (
	"ady-trans-jaya/config"
	"ady-trans-jaya/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetCars(c *gin.Context) {
	var cars []models.Cars
	if err := config.DB.Find(&cars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, cars)
}

func CreateCars(c *gin.Context) {
	var NewCars models.Cars
	if err := c.ShouldBindJSON(&NewCars); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if err := config.DB.Create(&NewCars).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "Cars successfully created",
		"cars":    NewCars,
	})
}
