package domain

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	"github.com/ankitsingh10194/banking/errs"
	"github.com/ankitsingh10194/banking/logger"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type CustomerRepositoryDb struct {
	client  *sql.DB
	client1 *sqlx.DB
}

func (c CustomerRepositoryDb) FindAll() ([]Customer, *errs.AppError) {
	sql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
	rows, err := c.client.Query(sql)
	if err != nil {
		logger.Error("Error while querying the customer table " + err.Error())
		return nil, errs.NewUnexpectedError("Database unexpected error")
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		var c Customer
		err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
		if err != nil {
			logger.Error("Error while scannning the customers: " + err.Error())
			return nil, errs.NewUnexpectedError("Database unexpected error")
		}
		customers = append(customers, c)
	}

	return customers, nil
}

func (c CustomerRepositoryDb) FindAllByStatus(status string) ([]Customer, *errs.AppError) {
	var err error
	cust := make([]Customer, 0)
	if status == "" {
		sql := "select customer_id,name,city,zipcode,date_of_birth,status from customers"
		err = c.client1.Select(&cust, sql)
	} else {
		sql := "select customer_id,name,city,zipcode,date_of_birth,status from customers where status=?"
		err = c.client1.Select(&cust, sql, status)
	}

	if err != nil {
		logger.Error("Error while querying customers table " + err.Error())
		return nil, errs.NewUnexpectedError("Database unexpected error")
	}
	// rows, err := c.client.Query(sql, status)
	// if err != nil {
	// 	log.Println("Error while querying the customer table", err.Error())
	// 	return nil, errs.NewUnexpectedError("Unexpected database error")
	// }
	// cust := make([]Customer, 0)
	// for rows.Next() {
	// 	var c Customer
	// 	err := rows.Scan(&c.Id, &c.Name, &c.City, &c.ZipCode, &c.DateOfBirth, &c.Status)
	// 	if err != nil {
	// 		return nil, errs.NewUnexpectedError("Unexpected database error")
	// 	}
	// 	cust = append(cust, c)
	// }
	return cust, nil
}

func (c CustomerRepositoryDb) GetCustomer(customerId string) (*Customer, *errs.AppError) {
	query := "select customer_id,name,city,zipcode,date_of_birth,status from customers where customer_id=?"
	var cust Customer
	err := c.client1.Get(&cust, query, customerId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			logger.Error("Error while scanning the customers" + err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}
	/*row := c.client.QueryRow(query, customerId)

	var cust Customer
	err := row.Scan(&cust.Id, &cust.Name, &cust.City, &cust.ZipCode, &cust.DateOfBirth, &cust.Status)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errs.NewNotFoundError("Customer not found")
		} else {
			log.Println("Error while scanning the customers", err.Error())
			return nil, errs.NewUnexpectedError("Unexpected database error")
		}
	}*/
	return &cust, nil
}

func NewCustomerRepositoryDb() CustomerRepositoryDb {
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbAddr := os.Getenv("DB_ADDR")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")
	dataSource := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUser, dbPassword, dbAddr, dbPort, dbName)

	client, err := sql.Open("mysql", dataSource)
	if err != nil {
		panic(err)
	}
	// See "Important settings" section.
	client.SetConnMaxLifetime(time.Minute * 3)
	client.SetMaxOpenConns(10)
	client.SetMaxIdleConns(10)

	return CustomerRepositoryDb{client: client}
}

func NewCustomerRepositoryDbSqlx(client *sqlx.DB) CustomerRepositoryDb {
	return CustomerRepositoryDb{client1: client}
}
