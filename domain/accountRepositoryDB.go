package domain

import (
	"strconv"

	"github.com/ankitsingh10194/banking/errs"
	"github.com/ankitsingh10194/banking/logger"
	"github.com/jmoiron/sqlx"
)

type AccountRepositoryDb struct {
	client *sqlx.DB
}

func (a AccountRepositoryDb) Save(acc Account) (*Account, *errs.AppError) {
	sqlInsert := "INSERT INTO accounts(customer_id,opening_date,account_type,amount,status) values(?,?,?,?,?)"
	res, err := a.client.Exec(sqlInsert, acc.CustomerId, acc.OpeningDate, acc.AccountType, acc.Amount, acc.Status)
	if err != nil {
		logger.Error("Error occurred while creating the new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}

	id, err := res.LastInsertId()
	if err != nil {
		logger.Error("Error occurred while getting last insert id for new account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected error from database")
	}
	acc.AccountId = strconv.FormatInt(id, 10)
	return &acc, nil
}

func (a AccountRepositoryDb) SaveTransaction(t Transaction) (*Transaction, *errs.AppError) {
	tx, err := a.client.Begin()
	if err != nil {
		logger.Error("Error while creating the new transaction for bank account transaction" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	// inserting bank account transaction
	insertQuery := "insert into transactions(account_id,amount,transaction_type,transaction_date) values(?,?,?,?)"
	result, err1 := tx.Exec(insertQuery, t.AccountId, t.Amount, t.TransactionType, t.TransactionDate)
	if err1 != nil {
		logger.Error("Unexpected database error" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}

	//update account balance
	if t.IsWithdrawal() {
		_, err = tx.Exec(`update accounts SET amount=amount- ? where account_id=?`, t.Amount, t.AccountId)
	} else {
		_, err = tx.Exec(`update accounts SET amount=amount+ ? where account_id=?`, t.Amount, t.AccountId)

	}
	// in case of error Rollback, and changes from both the tables will be reverted
	if err != nil {
		tx.Rollback()
		logger.Error("Error while saving the transaction for bank account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error!!")
	}
	// commit the transaction if all is good
	err = tx.Commit()
	if err != nil {
		tx.Rollback()
		logger.Error("Error while commit the transaction for bank account" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error!!")
	}

	//getting the last transaction Id from the transaction table
	transactionId, err2 := result.LastInsertId()
	if err2 != nil {
		logger.Error("Error while fetching the last transaction id" + err2.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	//fetch the latest account details from account table
	account, appErr := a.FindBy(t.AccountId)
	if appErr != nil {
		return nil, appErr
	}
	t.TransactionId = strconv.FormatInt(transactionId, 10)
	t.Amount = account.Amount
	return &t, nil
}

func (a AccountRepositoryDb) FindBy(accountId string) (*Account, *errs.AppError) {
	var acc Account
	sqlQuery := "SELECT account_id,customer_id,opening_date,account_type,amount from accounts where account_id=?"
	err := a.client.Get(&acc, sqlQuery, accountId)
	if err != nil {
		logger.Error("Error while fetching the account information" + err.Error())
		return nil, errs.NewUnexpectedError("Unexpected database error")
	}
	return &acc, nil

}
func NewAccountRepositoryDb(client *sqlx.DB) AccountRepositoryDb {
	return AccountRepositoryDb{client: client}

}
