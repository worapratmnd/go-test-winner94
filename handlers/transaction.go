package handlers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"main/model"
	"main/repository"
	"main/services"

	"github.com/gin-gonic/gin"
)

type handler struct {
	repo    repository.WinnerRepo
	service services.WinnerService
}

type Handler interface {
	GetTransaction(c *gin.Context)
	UpdateTransaction(c *gin.Context)
}

func NewHandler(transactionRepo repository.WinnerRepo, transService services.WinnerService) Handler {
	return &handler{
		repo:    transactionRepo,
		service: transService,
	}
}

func (p *handler) GetTransaction(c *gin.Context) {
	trans, err := p.service.GetTransactionList()
	if err != nil {
		fmt.Println("GetTransactionList err: ", err.Error())
		c.JSON(500, "error")
	} else {
		c.JSON(200, trans)
	}
}

func (p *handler) UpdateTransaction(c *gin.Context) {
	var (
		transReq model.TransactionDb
	)
	bodyReq, err := ioutil.ReadAll(c.Request.Body)
	if err = json.Unmarshal(bodyReq, &transReq); err != nil {
		fmt.Println(err.Error())
		c.JSON(500, "Error")
		return
	}

	trans, err := p.service.UpdateTransaction(transReq)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(500, "Error")
	} else {
		c.JSON(200, trans)
	}
}
