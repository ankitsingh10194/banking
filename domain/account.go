package domain

import (
	"github.com/ankit/banking/dto"
	"github.com/ankit/banking/errs"
)

type Account struct {
	AccountId   string  `db:"account_id"`
	CustomerId  string  `db:"customer_id"`
	OpeningDate string  `db:"opening_date"`
	AccountType string  `db:"account_type"`
	Amount      float64 `db:"amount"`
	Status      string  `db:"status"`
}

func (a Account) ToDto() dto.AccountResponse {
	res := dto.AccountResponse{AccountId: a.AccountId}
	return res
}

type AccountRepository interface {
	Save(Account) (*Account, *errs.AppError)
	SaveTransaction(Transaction) (*Transaction, *errs.AppError)
	FindBy(accountId string) (*Account, *errs.AppError)
}

func (a Account) CanWithdrawal(amount float64) bool {
	return a.Amount < amount
}
