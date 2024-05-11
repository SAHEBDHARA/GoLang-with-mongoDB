package helper

import (
	"context"
	model "dbconnect/models"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func AddTransaction(collection *mongo.Collection, transaction model.Transaction) {
	inserted, err := collection.InsertOne(context.Background(), transaction)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted 1 transaction in db with id:", inserted.InsertedID)
}

func UpdateTransaction(collection *mongo.Collection, transactionID string) {
	id, _ := primitive.ObjectIDFromHex(transactionID)
	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"isPaied": true}}
	result, err := collection.UpdateByID(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Transaction update successful ", result.UpsertedID)
}


func GetAllTransaction(collection *mongo.Collection)[]primitive.M{
	cursor, err := collection.Find(context.Background(),bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	var Transactions []primitive.M

	for cursor.Next(context.Background()){
		var Transaction bson.M
		err := cursor.Decode(&Transaction)
		if err != nil {
			log.Fatal(err)
		}
		Transactions = append(Transactions, Transaction)
	}
	defer cursor.Close(context.Background())
	return Transactions
}