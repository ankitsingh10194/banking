package service

import (
	"time"

	"github.com/ankit/banking/domain"
	"github.com/ankit/banking/dto"
	"github.com/ankit/banking/errs"
)

const dbTSLayout = "2006-01-02 15:04:05"

type AccountService interface {
	NewAccount(dto.AccountRequest) (*dto.AccountResponse, *errs.AppError)
	MakeTransaction(dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError)
}

type AccountServiceImpl struct {
	repo domain.AccountRepository
}

func (a AccountServiceImpl) NewAccount(acc dto.AccountRequest) (*dto.AccountResponse, *errs.AppError) {
	err := acc.Validate()
	if err != nil {
		return nil, err
	}
	//conversion dto to domain object
	account := domain.Account{
		AccountId:   "",
		CustomerId:  acc.CustomerId,
		OpeningDate: time.Now().Format("2006-01-02 15:04:05"),
		AccountType: acc.AccountType,
		Amount:      acc.Amount,
		Status:      "1",
	}
	newAccount, err := a.repo.Save(account)
	if err != nil {
		return nil, err
	}
	//conversion domain to dto
	res := newAccount.ToDto()

	return &res, nil
}

func (a AccountServiceImpl) MakeTransaction(req dto.TransactionRequest) (*dto.TransactionResponse, *errs.AppError) {
	err := req.Validate()
	if err != nil {
		return nil, err
	}
	if req.IsTransactionTypeWithdrawal() {
		account, err := a.repo.FindBy(req.AccountId)
		if err != nil {
			return nil, err
		}
		if account.CanWithdrawal(req.Amount) {
			return nil, errs.NewValidationError("Insufficient balance in the account")

		}

	}

	d := domain.Transaction{
		AccountId:       req.AccountId,
		TransactionType: req.TransactionType,
		Amount:          req.Amount,
		TransactionDate: time.Now().Format(dbTSLayout),
	}
	transaction, err := a.repo.SaveTransaction(d)
	if err != nil {
		return nil, err
	}

	response := transaction.ToDto()
	return &response, nil
}

func NewAccountService(repo domain.AccountRepository) AccountServiceImpl {
	return AccountServiceImpl{repo: repo}
}
