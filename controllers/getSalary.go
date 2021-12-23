package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/thonguyen/rest-api-golang/routes"

	"go.mongodb.org/mongo-driver/bson"
)

func GetPeopleBySalary(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	// Get average salary all person in one month
	// clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// client, _ := mongo.Connect(context.TODO(), clientOptions)
	// c := client.Database("example").Collection("people")
	fromDate := time.Date(2021, time.May, 1, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.June, 1, 23, 59, 59, 0, time.UTC)
	pipe := []bson.M{
		{"$match": bson.M{
			"dob": bson.M{
				"$gte": fromDate,
				"$lt":  toDate,
			},
		}},
		{"$group": bson.M{
			"_id":   "",
			"count": bson.M{"$sum": 1},
			"sum":   bson.M{"$sum": "$salary"},
			"avg":   bson.M{"$avg": "$salary"},
		}},
	}
	cursor, err := routes.ConnectDB().Aggregate(context.TODO(), pipe)
	if err != nil {
		panic(err)
	}
	var results []bson.M
	if err = cursor.All(context.TODO(), &results); err != nil {
		panic(err)
	}
	if err := cursor.Close(context.TODO()); err != nil {
		panic(err)
	}
	json.NewEncoder(response).Encode(results)
}
