package router

import (
	"dbconnect/controller"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message":"server is up and runnig.."})
	})
	router.GET("/transactions",controller.Alltransaction)
	router.POST("/transaction",controller.AddTransactionController)
	router.PUT("/transaction/:id", controller.UpdateTransaction)

	router.Run(":8080")
}