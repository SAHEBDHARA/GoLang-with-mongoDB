package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Transaction struct {
	ID      primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	User    string             `json:"user,omitempty`
	Ammount int                `json:"ammount,omitempty"`
	IsPaied bool               `json:"isPaied,omitempty`
}
