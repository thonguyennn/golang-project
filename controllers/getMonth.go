package controllers

import (
	"context"
	"encoding/json"
	"net/http"
	"time"

	"github.com/thonguyen/rest-api-golang/models"
	"github.com/thonguyen/rest-api-golang/routes"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPeopleByMonth(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []models.Person
	fromDate := time.Date(2021, time.May, 1, 12, 00, 0, 0, time.UTC) // year, month, day, hour, minute, second
	toDate := time.Date(2021, time.June, 1, 13, 00, 0, 0, time.UTC)  // year, month, day, hour, minute, second
	filter := bson.M{
		"dob": bson.M{
			"$gte": primitive.NewDateTimeFromTime(fromDate),
			"$lte": primitive.NewDateTimeFromTime(toDate),
		},
	}
	cursor, err := routes.ConnectDB().Find(context.TODO(), filter)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()) {
		var person models.Person
		cursor.Decode(&person)
		people = append(people, person)
	}
	if err := cursor.Err(); err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(people)
}
