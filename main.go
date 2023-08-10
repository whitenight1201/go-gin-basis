package main

import (
	routes "CRUD_API/routes"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Setup route group for the API
	api := router.Group("/api")
	{
		api.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"message": "hello",
			})
		})
	}

	api.POST("/", routes.CreatePost)

	// called as localhost:3000/getOne/{id}
	api.GET("getOne/:postId", routes.ReadOnePost)

	// called as localhost:3000/update/{id}
	api.PUT("/update/:postId", routes.UpdatePost)

	// called as localhost:3000/delete/{id}
	api.DELETE("/delete/:postId", routes.DeletePost)

	router.Run("localhost: 3000")
}
