package controllers

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"net/http"

	"github.com/thonguyen/rest-api-golang/cache"
	"github.com/thonguyen/rest-api-golang/models"
	"github.com/thonguyen/rest-api-golang/routes"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreatePeople(response http.ResponseWriter, request *http.Request) {
	var person models.Person
	response.Header().Set("content-type", "application/json")
	json.NewDecoder(request.Body).Decode(&person)
	result, _ := routes.ConnectDB().InsertOne(context.TODO(), person)
	json.NewEncoder(response).Encode(result)
}

func GetPeopleById(response http.ResponseWriter, request *http.Request) {
	// get redis database
	redisDB := cache.GetRedisDB()

	response.Header().Set("content-type", "application/json")
	var person models.Person

	id, _ := primitive.ObjectIDFromHex(mux.Vars(request)["id"])
	idUser := id.Hex()

	value, err := redisDB.Get(context.Background(), idUser).Result()
	if err != nil {
		fmt.Println("Get data from mongodb")
		result := routes.ConnectDB().FindOne(context.TODO(), bson.M{"_id": id}).Decode(&person)
		if result != nil {
			response.WriteHeader(http.StatusInternalServerError)
			response.Write([]byte(`{ "message": "` + result.Error() + `" }`))
			return
		}
		json.NewEncoder(response).Encode(person)
		a, _ := json.Marshal(person)
		redisDB.Set(context.Background(), idUser, a, 0).Err()
	} else {
		fmt.Println("Get data from redis")
		json.NewDecoder(strings.NewReader(value)).Decode(&person)
		json.NewEncoder(response).Encode(person)
	}
}

func GetPeople(response http.ResponseWriter, request *http.Request) {
	// get redis database
	redisDB := cache.GetRedisDB()

	response.Header().Set("content-type", "application/json")
	var people []models.Person

	value, err := redisDB.Get(context.Background(), "people").Result()
	if err != nil {
		fmt.Println("Get data from mongodb")
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
		a, _ := json.Marshal(people)
		redisDB.Set(context.Background(), "people", a, 0).Err()
	} else {
		fmt.Println("Get data from redis")
		json.NewDecoder(strings.NewReader(value)).Decode(&people)
		json.NewEncoder(response).Encode(people)
	}

}

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
