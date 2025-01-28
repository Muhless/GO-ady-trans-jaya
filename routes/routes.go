package routes

import (
	"ady-trans-jaya/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.SetTrustedProxies(nil)

	r.GET("/api/users", controllers.GetUsers)
	r.GET("/api/users/:id", controllers.GetUserByID)
	r.POST("/api/users", controllers.CreateUser)
	r.PUT("/api/users/:id", controllers.UpdateUser)
	r.DELETE("/api/users/:id", controllers.DeleteUser)

	r.GET("/api/cars", controllers.GetCars)
	// r.GET("/api/cars/:id", controllers.GetCarByID)
	r.POST("/api/cars", controllers.CreateCars)
	// r.PUT("/api/cars/:id", controllers.UpdateCar)
	// r.DELETE("/api/cars/:id", controllers.DeleteCar)

	return r
}
