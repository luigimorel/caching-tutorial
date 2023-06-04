package routes

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/morelmiles/go-redis-caching/internals/controllers"
	"github.com/morelmiles/go-redis-caching/internals/helpers"
)

func Routes() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()
	v1 := r.Group("/v1")

	// User routes
	v1.GET("/users", controllers.GetUsers)
	v1.GET("/users/:id", controllers.GetUserById)
	v1.PUT("/users/:id", controllers.UpdateUserById)
	v1.DELETE("/users/:id", controllers.DeleteUserById)
	v1.POST("/register", controllers.CreateUser)

	r.Use(helpers.RequestTimeout)

	r.Run(":" + os.Getenv("PORT"))
}
