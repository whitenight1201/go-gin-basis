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

func ReadOnePost(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	var DB = database.ConnectDB()
	var postCollection = getcollection.GetCollection(DB, "Posts")

	//The postId variable is a parameter declaration. It gets a document's object ID as objId.
	//However, result is an instance of the database model, which later holds the returned document as res.
	postId := c.Param("postId")
	var result model.Posts

	defer cancel()

	objId, _ := primitive.ObjectIDFromHex(postId)

	err := postCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&result)

	res := map[string]interface{}{"data": result}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "success!", "Data": res})
}
