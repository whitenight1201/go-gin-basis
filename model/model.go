// This handles database model
package model

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Posts struct {
	ID      primitive.ObjectID
	Title   string
	Article string
}
