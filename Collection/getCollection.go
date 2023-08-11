package getcollection

import (
	"go.mongodb.org/mongo-driver/mongo"
)

// This function gets the Collection from the MongoDB database.
// The database name, in this case, is myGoappDB, with Posts as its collection.
func GetCollection(client *mongo.Client, collectionName string) *mongo.Collection {
	collection := client.Database("myGoappDB").Collection(collectionName)
	return collection
}
