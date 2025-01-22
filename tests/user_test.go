package tests

import (
	"ady-trans-jaya/config"
	"ady-trans-jaya/controllers"
	"ady-trans-jaya/models"
	"bytes"
	"encoding/json"

	// "ady-trans-jaya/models"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	router := gin.Default()
	return router
}

func resetDatabase() {
	config.DB.Exec("DELETE FROM users")
}

func TestGetUsers(t *testing.T) {
	config.ConnectDB()
	resetDatabase()
	config.DB.Create(&models.User{Username: "testuser", Password: "password"})

	router := setupRouter()
	router.GET("/users", controllers.GetUsers)

	req, _ := http.NewRequest("GET", "/users", nil)
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
}

func TestCreateUser(t *testing.T) {
	config.ConnectDB()
	resetDatabase()
	router := setupRouter()
	router.POST("/users", controllers.CreateUser)

	payload := models.User{
		Username: "testingcreate",
		Password: "testingcreate123",
	}
	body, _ := json.Marshal(payload)

	req, _ := http.NewRequest("POST", "/users", bytes.NewBuffer(body))
	req.Header.Set("Content-Type", "application-json")
	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)

	var response map[string]interface{}
	err := json.Unmarshal(w.Body.Bytes(), &response)
	assert.NoError(t, err)
	assert.Equal(t, "User successfully created", response["message"])

	var CreateUser models.User
	err = config.DB.Where("username = ?", payload.Username).First(&CreateUser).Error
	assert.NoError(t, err)
	assert.Equal(t, payload.Username, CreateUser.Username)
}
