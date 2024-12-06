package db

import (
    "context"
    "crypto/tls"
    "log"
    "os"
    "time"

    "github.com/joho/godotenv"
    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

var QuotesCollection *mongo.Collection

func ConnectMongoDB() {
    // Load the .env file
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }

    // Get the MongoDB URI from the .env file
    mongoURI := os.Getenv("MONGO_URI")
    if mongoURI == "" {
        log.Fatalf("MONGO_URI not found in .env file")
    }

    // Load the client certificate
    clientCertFile := "certs/client.pem"
    clientCert, err := tls.LoadX509KeyPair(clientCertFile, clientCertFile)
    if err != nil {
        log.Fatalf("Failed to load client certificate: %v", err)
    }

    // Create TLS config
    tlsConfig := &tls.Config{
        Certificates:       []tls.Certificate{clientCert},
        InsecureSkipVerify: false, // Don't skip verification in production
    }

    // MongoDB client options
    clientOpts := options.Client().
        ApplyURI(mongoURI).
        SetTLSConfig(tlsConfig)

    // Connect to MongoDB
    client, err := mongo.NewClient(clientOpts)
    if err != nil {
        log.Fatalf("Failed to create MongoDB client: %v", err)
    }

    ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
    defer cancel()

    err = client.Connect(ctx)
    if err != nil {
        log.Fatalf("Failed to connect to MongoDB: %v", err)
    }

    // Ping the database to verify the connection
    err = client.Ping(ctx, nil)
    if err != nil {
        log.Fatalf("Failed to ping MongoDB: %v", err)
    }

    log.Println("Connected to MongoDB successfully")

    // Assign the Quotes collection
    QuotesCollection = client.Database("HerHourlyLiftDB").Collection("quotes")
}
