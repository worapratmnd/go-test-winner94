package handlers_test

import (
	"bytes"
	"encoding/json"
	"errors"
	"main/handlers"
	"main/model"
	"main/repository"
	"main/services"
	"main/utils"
	. "main/utils"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestGetTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		transactionRepo := repository.NewWinnerDbMock()
		transService := services.NewWinnerService(transactionRepo)

		handlers := handlers.NewHandler(transactionRepo, transService)
		gin.SetMode(gin.TestMode)
		route := gin.Default()
		route.GET("/", handlers.GetTransaction)
		transactionRepo.On("GetTransactionList").Return([]model.TransactionDb{
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

		// Act
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Error", func(t *testing.T) {
		transactionRepo := repository.NewWinnerDbMock()
		transService := services.NewWinnerService(transactionRepo)
		handlers := handlers.NewHandler(transactionRepo, transService)

		gin.SetMode(gin.TestMode)
		route := gin.Default()
		route.GET("/", handlers.GetTransaction)
		transactionRepo.On("GetTransactionList").Return([]model.TransactionDb{}, errors.New("Test"))

		// Act
		req := httptest.NewRequest("GET", "/", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
}

func TestUpdateTransaction(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		var (
			resMock  int64
			tranMock model.TransactionDb
		)
		transactionRepo := repository.NewWinnerDbMock()
		transService := services.NewWinnerService(transactionRepo)
		handlers := handlers.NewHandler(transactionRepo, transService)

		gin.SetMode(gin.TestMode)
		route := gin.Default()
		route.POST("/update", handlers.UpdateTransaction)

		resMock = 1
		tranMock = model.TransactionDb{
			Id:    111,
			Input: utils.ReturnStringPointer("124+600"),
		}
		transactionRepo.On("UpdateTransaction").Return(resMock, nil)

		// Act
		jsonValue, _ := json.Marshal(tranMock)
		req := httptest.NewRequest("POST", "/update", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
		assert.Equal(t, http.StatusOK, w.Code)
	})

	t.Run("Error Empty", func(t *testing.T) {
		transactionRepo := repository.NewWinnerDbMock()
		transService := services.NewWinnerService(transactionRepo)

		handlers := handlers.NewHandler(transactionRepo, transService)

		gin.SetMode(gin.TestMode)
		route := gin.Default()
		route.POST("/update", handlers.UpdateTransaction)

		// Act
		req := httptest.NewRequest("POST", "/update", nil)
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})

	t.Run("Error", func(t *testing.T) {
		var (
			resMock  int64
			tranMock model.TransactionDb
		)

		transactionRepo := repository.NewWinnerDbMock()
		transService := services.NewWinnerService(transactionRepo)

		handlers := handlers.NewHandler(transactionRepo, transService)

		gin.SetMode(gin.TestMode)
		route := gin.Default()
		route.POST("/update", handlers.UpdateTransaction)

		resMock = 0
		tranMock = model.TransactionDb{
			Id:    111,
			Input: utils.ReturnStringPointer("124+600"),
		}
		transactionRepo.On("UpdateTransaction").Return(resMock, errors.New("Test Error"))

		// Act
		jsonValue, _ := json.Marshal(tranMock)
		req := httptest.NewRequest("POST", "/update", bytes.NewBuffer(jsonValue))
		w := httptest.NewRecorder()
		route.ServeHTTP(w, req)
		assert.Equal(t, 500, w.Code)
	})
}
