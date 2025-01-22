package tests

import (
	"ady-trans-jaya/config"
	"ady-trans-jaya/controllers"
	"ady-trans-jaya/models"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func init() {
	config.ConnectDB()
}

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/users", controllers.GetUsers)
	r.GET("/users/:id", controllers.GetUserByID)
	r.POST("/users", controllers.CreateUser)
	r.PUT("/users/:id", controllers.UpdateUser)
	r.DELETE("/users/:id", controllers.DeleteUser)

	return r
}

func TestGetUsers(t *testing.T) {
	router := SetupRouter()

	//seeding
	config.DB.Create(&models.User{Username: "muhless", Password: "muhless123"})
	config.DB.Create(&models.User{Username: "onana", Password: "onana123"})

	// create request to endpoint
	req, _ := http.NewRequest("GET", "/users", nil)
	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var users []models.User
	err := json.Unmarshal(resp.Body.Bytes(), &users)
	assert.Nil(t, err)
	assert.GreaterOrEqual(t, len(users), 2)
}

// func TestGetUsersByID(t *testing.T) {
// 	router := gin.Default()
// 	router.GET("/users/:id", controllers.GetUserByID())
// }
