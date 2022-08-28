package services

import (
	"main/model"
	"main/repository"
)

type winnerService struct {
	winnerRepo repository.WinnerRepo
}

type WinnerService interface {
	GetTransactionList() ([]model.TransactionDb, error)
	UpdateTransaction(transItem model.TransactionDb) (int64, error)
}

func NewWinnerService(repo repository.WinnerRepo) WinnerService {
	return &winnerService{winnerRepo: repo}
}

func (p *winnerService) GetTransactionList() ([]model.TransactionDb, error) {
	resModel, resError := p.winnerRepo.GetTransactionList()

	return resModel, resError
}

func (p *winnerService) UpdateTransaction(transItem model.TransactionDb) (int64, error) {
	res, resError := p.winnerRepo.UpdateTransaction(transItem)

	return res, resError
}
