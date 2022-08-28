package services_test

import (
	"fmt"
	"main/model"
	"main/repository"
	"main/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetTransactionList(t *testing.T) {
	transRepo := repository.NewWinnerDbMock()
	transRepo.On("GetTransactionList").Return([]model.TransactionDb{
		{
			Id:         1,
			Input:      ReturnStringPointer("111+600"),
			Type:       ReturnStringPointer("TRANSACTION"),
			S1:         ReturnIntPointer(1),
			S2:         ReturnIntPointer(1),
			S3:         ReturnIntPointer(1),
			Total:      ReturnStringPointer("600"),
			Status:     ReturnIntPointer(1),
			CreateBy:   ReturnStringPointer("giftsupport"),
			CreateDttm: ReturnStringPointer("2022-08-06 11:11:53"),
			UpdateBy:   ReturnStringPointer("giftsupport"),
			UpdateDttm: ReturnStringPointer("2022-08-06 11:11:53"),
			Tables:     ReturnIntPointer(1),
			RoundName:  ReturnStringPointer("2022-08-06"),
			RoundNo:    ReturnIntPointer(1),
			Result:     ReturnStringPointer("WIN"),
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
		Input: ReturnStringPointer("111+6000"),
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

func ReturnStringPointer(s string) *string {
	v := s
	return &v
}

func ReturnIntPointer(s int) *int {
	v := s
	return &v
}
