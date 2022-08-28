package repository

import (
	"main/model"

	"github.com/stretchr/testify/mock"
)

type WinnerRepoMock struct {
	mock.Mock
}

func NewWinnerDbMock() *WinnerRepoMock {
	return &WinnerRepoMock{}
}

func (m *WinnerRepoMock) GetTransactionList() ([]model.TransactionDb, error) {
	args := m.Called()
	return args.Get(0).([]model.TransactionDb), args.Error(1)
}

func (m *WinnerRepoMock) UpdateTransaction(transItem model.TransactionDb) (int64, error) {
	args := m.Called()
	return args.Get(0).(int64), args.Error(1)
}
