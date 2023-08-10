package routes

import (
	getcollection "CRUD_API/Collection"
	database "CRUD_API/database"
	model "CRUD_API/model"
	"context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func UpdatePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "Posts")

	postId := c.Param("postId")
	var post model.Posts

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	if err := c.BindJSON(&post); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	//A JSON format of the model instance (post) calls each model field from the database.
	edited := bson.M{"title": post.Title, "article": post.Article}

	//The result variable uses the MongoDB $set operator to update a required document called by its object ID.
	result, err := postCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": edited})

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	//The result.MatchedCount condition prevents the code from running if there's no record in the database or the passed ID is invalid.
	if result.MatchedCount < 1 {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Data doesn't exist"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "data updated successfully!", "Data": res})
}
