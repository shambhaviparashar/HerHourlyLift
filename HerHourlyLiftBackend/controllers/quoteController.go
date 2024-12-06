package controllers

import (
	"context"
	"log"
	"time"

	"HerHourlyLiftBackend/db"
	"HerHourlyLiftBackend/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// GetRandomQuote retrieves a random quote from MongoDB
func GetRandomQuote() (models.Quote, error) {
	var quote models.Quote
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// MongoDB aggregation to fetch a random quote
	pipeline := mongo.Pipeline{
		{{"$sample", bson.D{{"size", 1}}}}, // Random sampling
	}

	cursor, err := db.QuotesCollection.Aggregate(ctx, pipeline, options.Aggregate())
	if err != nil {
		log.Printf("Error during aggregation: %v", err)
		return quote, err
	}
	defer cursor.Close(ctx)

	// Check if the cursor has a document
	if cursor.Next(ctx) {
		if err := cursor.Decode(&quote); err != nil {
			log.Printf("Error decoding quote: %v", err)
			return quote, err
		}
	} else {
		log.Println("No quotes found in the database")
		return quote, mongo.ErrNoDocuments
	}

	return quote, nil
}
