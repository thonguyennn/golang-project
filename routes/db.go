package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection *mongo.Collection
var DB *mongo.Database

func ConnectDB() *mongo.Collection {
	// config := GetConfiguration()

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")

	// Connect to mongoDB

	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	collection := client.Database("example").Collection("people")

	return collection
}

type ErrorResponse struct {
	StatusCode   int    `json:"status"`
	ErrorMessage string `json:"massage"`
}

func GetError(err error, w http.ResponseWriter) {
	log.Fatal(err.Error())
	var response = ErrorResponse{
		ErrorMessage: err.Error(),
		StatusCode:   http.StatusInternalServerError,
	}

	message, _ := json.Marshal(response)

	w.WriteHeader(response.StatusCode)
	w.Write(message)
}

// Configuration model
type Configuration struct {
	Port             string
	ConnectionString string
}
