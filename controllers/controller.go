package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	// "time"

	"github.com/thonguyen/rest-api-golang/models"
	"github.com/thonguyen/rest-api-golang/routes"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	// "go.mongodb.org/mongo-driver/mongo"
	// "go.mongodb.org/mongo-driver/mongo/options"
)

type PeopleControllers struct{}

func CreatePeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person models.Person
	json.NewDecoder(request.Body).Decode(&person)
	result, _ := routes.ConnectDB().InsertOne(context.TODO(), person)
	json.NewEncoder(response).Encode(result)

}

func GetPeopleById(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person models.Person
	params := mux.Vars(request)
	id, _ := primitive.ObjectIDFromHex(params["id"])
	err := routes.ConnectDB().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&person)
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(person)
}

func GetPeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var people []models.Person
	cursor, err := routes.ConnectDB().Find(context.TODO(), bson.M{})
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

// func GetPeopleBySalary(response http.ResponseWriter, request *http.Request) {
// 	response.Header().Set("content-type", "application/json")
// 	// Get average salary all person in one month
// 	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
// 	client, _ := mongo.Connect(context.TODO(), clientOptions)
// 	c := client.Database("example").Collection("people")
// 	fromDate := time.Date(2021, time.May, 1, 10, 25, 35, 0, time.UTC)
// 	toDate := time.Date(2021, time.June, 1, 11, 50, 24, 352, time.UTC)
// 	pipe := []bson.M{
// 		{"$match": bson.M{
// 			"dob": bson.M{
// 				"$gte": fromDate,
// 				"$lt":  toDate,
// 			},
// 		}},
// 		{"$group": bson.M{
// 			"_id":   "",
// 			"count": bson.M{"$sum": 1},
// 			"sum":   bson.M{"$sum": "$salary"},
// 			"avg":   bson.M{"$avg": "$salary"},
// 		}},
// 	}
// 	cursor, err := c.Aggregate(context.TODO(), pipe)
// 	if err != nil {
// 		panic(err)
// 	}
// 	var results []bson.M
// 	if err = cursor.All(context.TODO(), &results); err != nil {
// 		panic(err)
// 	}
// 	if err := cursor.Close(context.TODO()); err != nil {
// 		panic(err)
// 	}
// 	json.NewEncoder(response).Encode(results)
// }

func UpdatePeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	var person models.Person
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	json.NewDecoder(request.Body).Decode(&person)
	result, err := routes.ConnectDB().UpdateOne(context.TODO(), models.Person{ID: id}, bson.M{"$set": person})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func DeletePeople(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("content-type", "application/json")
	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	result, err := routes.ConnectDB().DeleteOne(context.TODO(), bson.M{"_id": id})
	if err != nil {
		response.WriteHeader(http.StatusInternalServerError)
		response.Write([]byte(`{ "message": "` + err.Error() + `" }`))
		return
	}
	json.NewEncoder(response).Encode(result)
}

func HomeLink(response http.ResponseWriter, request *http.Request) {
	fmt.Fprintf(response, "Welcome home!")
}
