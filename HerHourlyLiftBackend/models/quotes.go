package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// Quote defines the structure of a quote document in MongoDB
type Quote struct {
	ID    primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Quote string             `json:"quote" bson:"quote"`
}
