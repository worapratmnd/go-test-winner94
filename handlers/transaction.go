package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/database"
	"main/model"
	"main/repository"
	"main/services"

	"github.com/gin-gonic/gin"
)

func GetTransaction(c *gin.Context) {
	transactionRepo := repository.NewWinnerDb(database.MyWinnerDB)
	transService := services.NewWinnerService(transactionRepo)
	trans, err := transService.GetTransactionList()
	if err != nil {
		fmt.Println("GetTransactionList err: ", err.Error())
	}

	c.JSON(200, trans)
}

func UpdateTransaction(c *gin.Context) {
	var (
		transReq model.TransactionDb
	)
	bodyReq, err := ioutil.ReadAll(c.Request.Body)
	if err = json.Unmarshal(bodyReq, &transReq); err != nil {
		fmt.Println(err.Error())
	}

	transactionRepo := repository.NewWinnerDb(database.MyWinnerDB)
	transService := services.NewWinnerService(transactionRepo)

	trans, err := transService.UpdateTransaction(transReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, "Error")
	} else {
		c.JSON(200, trans)
	}
}
