package dto

import (
	"strings"

	"github.com/ankit/banking/errs"
)

type TransactionRequest struct {
	CustomerId      string  `json:"customer_id"`
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"amount"`
	TransactionDate string  `json:"transaction_date"`
}

type TransactionResponse struct {
	TransactionId   string  `json:"transaction_id"`
	AccountId       string  `json:"account_id"`
	TransactionType string  `json:"transaction_type"`
	Amount          float64 `json:"new_balance"`
	TransactionDate string  `json:"transaction_date"`
}

func (t TransactionRequest) Validate() *errs.AppError {
	if !strings.EqualFold(t.TransactionType, "withdrawal") && !strings.EqualFold(t.TransactionType, "deposit") {
		return errs.NewValidationError("Transaction type can be only withdrawal or deposit!!")
	}
	if t.Amount < 0 {
		return errs.NewValidationError("Amount can not be less than zero!!")
	}
	return nil
}

func (t TransactionRequest) IsTransactionTypeWithdrawal() bool {
	return strings.EqualFold(t.TransactionType, "withdrawal")
}

func (t TransactionRequest) IsTransactionTypeDeposite() bool {
	return strings.EqualFold(t.TransactionType, "deposit")
}
