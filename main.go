package main

import (
	"main/database"
	"main/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	port := "7001"
	database.InitDB()

	gin.SetMode(gin.ReleaseMode)
	route := gin.Default()
	route.GET("/", handlers.GetTransaction)
	route.POST("/update", handlers.UpdateTransaction)
	route.Run(":" + port)
}
