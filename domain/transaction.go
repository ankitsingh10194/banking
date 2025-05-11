package domain

import (
	"strings"

	"github.com/ankit/banking/dto"
	"github.com/ankit/banking/errs"
)

type Transaction struct {
	TransactionId   string  `db:"transaction_id"`
	AccountId       string  `db:"account_id"`
	TransactionDate string  `db:"transaction_date"`
	TransactionType string  `db:"transaction_type"`
	Amount          float64 `db:"amount"`
}

type TransactionRepository interface {
	Deposite(Transaction) (*Transaction, *errs.AppError)
	Withdrawal(Transaction) (*Transaction, *errs.AppError)
}

func (t Transaction) IsWithdrawal() bool {
	return strings.EqualFold(t.TransactionType, "withdrawal")
}

func (t Transaction) ToDto() dto.TransactionResponse {
	return dto.TransactionResponse{
		TransactionId:   t.TransactionId,
		AccountId:       t.AccountId,
		TransactionType: t.TransactionType,
		Amount:          t.Amount,
		TransactionDate: t.TransactionDate,
	}
}
