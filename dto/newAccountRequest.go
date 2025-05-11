package dto

import (
	"strings"

	"github.com/ankit/banking/errs"
)

type AccountRequest struct {
	CustomerId  string  `json:"customer_id"`
	AccountType string  `json:"account_type"`
	Amount      float64 `json:"amount"`
}

func (req AccountRequest) Validate() *errs.AppError {
	if req.Amount < 5000 {
		return errs.NewValidationError("To open account, you need to deposit atleast 5000/-")
	}
	if strings.ToLower(req.AccountType) != "saving" && strings.ToLower(req.AccountType) != "checking" {
		return errs.NewValidationError("Account type should be checking or saving!!")
	}
	return nil
}
