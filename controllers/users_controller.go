package controllers

import (
	"ady-trans-jaya/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var users = []models.User{
	{ID: 1, Username: "Cihuy", Password: "cihuy12345"},
	{ID: 2,Username: "admin", Password: "admin123"},
}

func GetUsers(c *gin.Context) {
	c.JSON(http.StatusOK, users)
}

func GetUserByID(c *gin.Context) {
    id := c.Param("id")
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "invalid user ID"})
	}
    for _, user := range users {
        if user.ID == userID {
            c.JSON(http.StatusOK, user)
            return
        }
    }
    c.JSON(http.StatusNotFound, gin.H{"message": "User not found"})
}


func CreateUser(c *gin.Context)  {
	var newUser models.User
	if err := c.ShouldBindJSON(&newUser); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := newUser.HashPassword(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed hashing password"})
		return
	}

	users = append(users, newUser)
	c.JSON(http.StatusCreated, gin.H{
		"message": "User succesfully created",
		"user": newUser,
	})
}



