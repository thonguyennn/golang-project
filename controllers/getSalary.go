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
	fromDate := time.Date(2021, time.May, 1, 0, 0, 0, 0, time.UTC)
	toDate := time.Date(2021, time.June, 1, 0, 0, 0, 0, time.UTC)
	pipe := []bson.M{
		{"$match": bson.M{
			"dob": bson.M{
				"$gte": fromDate,
				"$lte": toDate,
			},
		}},
		{"$group": bson.M{
			// Trong 7 ngày
			// "_id": bson.M{
			// 	"month": bson.M{"$month": "$dob"},
			// 	"day":   bson.M{"$dayOfMonth": "$dob"},
			// },

			// Trong 1 ngày
			"_id": bson.M{
				"hour": bson.M{"$hour": "$dob"},
			},
			"count": bson.M{"$sum": 1},
			"avg":   bson.M{"$avg": "$salary"},
			"sum":   bson.M{"$sum": "$salary"},
		}},
		{"$sort": bson.M{
			"avg": 1,
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
