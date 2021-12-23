package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID          primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name        string             `json:"name,omitempty" bson:"name,omitempty"`
	Age         int                `json:"age,omitempty" bson:"age,omitempty"`
	Description string             `json:"description,omitempty" bson:"description,omitempty"`
	Address     string             `json:"address,omitempty" bson:"address,omitempty"`
	Salary      int                `json:"salary,omitempty" bson:"salary,omitempty"`
	Dob         time.Time          `json:"dob,omitempty" bson:"dob,omitempty"`
}
