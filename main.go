package main

import (
	"net/http"

	"github.com/Ryuuukin/ap-assignment1/controllers"
	"github.com/Ryuuukin/ap-assignment1/initializers"
	"github.com/Ryuuukin/ap-assignment1/middlewares"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVars()
	initializers.ConnectToDB()
}

func main() {
	r := gin.Default()

	r.Use(middlewares.RateLimitMiddleware)

	r.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status":  "success",
			"message": "Welcome to Pro E-Players Tracker!",
		})
	})

	r.POST("/users", controllers.UsersCreate)
	r.PUT("/users/:id", controllers.UsersUpdate)

	r.GET("/users", controllers.UsersIndex)
	r.GET("/users/:id", controllers.UsersShow)
	r.GET("users/filtered-sorted-paginated", controllers.FilteredUsersIndex)

	r.DELETE("/users/:id", controllers.UsersDelete)

	r.Run()
}
