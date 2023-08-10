package routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/database"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// This code deletes a record using the DeleteOne function.
func DeletePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	postId := c.Param("postId")

	var postCollection = getcollection.GetCollection(DB, "Posts")
	defer cancel()
	objId, _ := primitive.ObjectIDFromHex(postId)
	result, err := postCollection.DeleteOne(ctx, bson.M{"id": objId})
	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	//result.DeletedCount property prevents the code from running if the database is empty or the object ID is invalid.
	if result.DeletedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "No data to delete"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Article deleted successfully", "Data": res})
}
