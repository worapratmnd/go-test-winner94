package repository

import (
	"fmt"
	"main/model"

	"github.com/go-gorp/gorp"
)

type winnerDB struct {
	Db *gorp.DbMap
}

type WinnerRepo interface {
	GetTransactionList() ([]model.TransactionDb, error)
	UpdateTransaction(transItem model.TransactionDb) (int64, error)
}

func NewWinnerDb(db *gorp.DbMap) WinnerRepo {
	return &winnerDB{
		Db: db,
	}
}

func (p *winnerDB) GetTransactionList() ([]model.TransactionDb, error) {
	var (
		resModel []model.TransactionDb
		resError error
		queryStr string
	)

	queryStr = "select * from `transaction`"
	_, err := p.Db.Select(&resModel, queryStr)
	if err != nil {
		resError = err
	}

	return resModel, resError
}

func (p *winnerDB) UpdateTransaction(transItem model.TransactionDb) (int64, error) {
	var (
		resModel int64
		resError error
		queryStr string
	)

	// params := map[string]interface{}{}
	// params["input"] = transItem.Input
	// params["id"] = transItem.Id

	queryStr = " UPDATE `transaction` SET input = ? WHERE id = ?"

	result, err := p.Db.Exec(queryStr, transItem.Input, transItem.Id)
	if err != nil {
		resError = err
		fmt.Println(err.Error())
	} else {
		resModel, err = result.RowsAffected()
		if err != nil {
			fmt.Println(err.Error())
			resError = err
		}
	}
	return resModel, resError

}
