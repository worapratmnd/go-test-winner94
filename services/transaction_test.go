package services_test

import (
	"fmt"
	"main/model"
	"main/repository"
	"main/services"
	"main/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactionList(t *testing.T) {
	transRepo := repository.NewWinnerDbMock()
	transRepo.On("GetTransactionList").Return([]model.TransactionDb{
		{
			Id:         1,
			Input:      utils.ReturnStringPointer("111+600"),
			Type:       utils.ReturnStringPointer("TRANSACTION"),
			S1:         utils.ReturnIntPointer(1),
			S2:         utils.ReturnIntPointer(1),
			S3:         utils.ReturnIntPointer(1),
			Total:      utils.ReturnStringPointer("600"),
			Status:     utils.ReturnIntPointer(1),
			CreateBy:   utils.ReturnStringPointer("giftsupport"),
			CreateDttm: utils.ReturnStringPointer("2022-08-06 11:11:53"),
			UpdateBy:   utils.ReturnStringPointer("giftsupport"),
			UpdateDttm: utils.ReturnStringPointer("2022-08-06 11:11:53"),
			Tables:     utils.ReturnIntPointer(1),
			RoundName:  utils.ReturnStringPointer("2022-08-06"),
			RoundNo:    utils.ReturnIntPointer(1),
			Result:     utils.ReturnStringPointer("WIN"),
		},
	}, nil)

	transService := services.NewWinnerService(transRepo)
	res, err := transService.GetTransactionList()
	fmt.Printf("%+v", res)
	assert.NoError(t, err)
	assert.Equal(t, 1, res[0].Id)
}

func TestUpdateTransaction(t *testing.T) {
	var (
		mockRes   int64
		transMock model.TransactionDb
	)
	transMock = model.TransactionDb{
		Id:    887,
		Input: utils.ReturnStringPointer("111+6000"),
	}
	mockRes = 1

	transRepo := repository.NewWinnerDbMock()
	transService := services.NewWinnerService(transRepo)

	transRepo.On("UpdateTransaction").Return(mockRes, nil)
	res, err := transService.UpdateTransaction(transMock)
	fmt.Printf("%+v", res)
	assert.NoError(t, err)
	// assert.Equal(t, 1, res)
}
