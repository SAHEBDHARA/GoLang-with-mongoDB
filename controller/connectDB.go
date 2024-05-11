package controller

import (
	"context"
	helper "dbconnect/helpers"
	model "dbconnect/models"
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

const conncetionString = "mongodb://localhost:27017"
const dbName = "Bank"
const colName = "transaction"

var Collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(conncetionString)

	client, err := mongo.Connect(context.TODO(), clientOption)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("mongodb is connected")

	Collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is ready")
}

// add helpers

func addTransactionHelper(transaction model.Transaction) {
	helper.AddTransaction(Collection, transaction)
}

func updateTransactionHelper(transactionID string) {
	helper.UpdateTransaction(Collection, transactionID)
}

func getTransactionHelper() []primitive.M {
	return helper.GetAllTransaction(Collection)
}

// add controller

func Alltransaction(c *gin.Context) {
	transactions := getTransactionHelper()
	c.JSON(http.StatusOK, gin.H{"transaction":transactions})
}

func AddTransactionController(c *gin.Context) {
	var Transaction model.Transaction
	if err := c.ShouldBindJSON(&Transaction); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}
	addTransactionHelper(Transaction)
	c.Status(http.StatusCreated)
}

func UpdateTransaction(c *gin.Context) {
	transactionID := c.Param("id")
	updateTransactionHelper(transactionID)

	c.JSON(http.StatusOK, gin.H{"message": "Transaction updated successfully"})
}
