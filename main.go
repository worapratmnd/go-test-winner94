package main

import (
	"main/database"
	"main/handlers"
	"main/repository"
	"main/services"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "7001"
	database.InitDB()

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()

	transactionRepo := repository.NewWinnerDb(database.MyWinnerDB)
	transService := services.NewWinnerService(transactionRepo)
	handlers := handlers.NewHandler(transactionRepo, transService)
	route.GET("/", handlers.GetTransaction)
	route.POST("/update", handlers.UpdateTransaction)
	route.Run(":" + port)
}
