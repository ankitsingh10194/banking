package domain

import (
	"github.com/ankitsingh10194/banking/dto"
	"github.com/ankitsingh10194/banking/errs"
)

type Customer struct {
	Id          string `db:"customer_id"`
	Name        string
	City        string
	ZipCode     string
	DateOfBirth string `db:"date_of_birth"`
	Status      string
}

func (c Customer) statusToText() string {
	statusToText := "active"
	if c.Status == "0" {
		statusToText = "inActive"
	}
	return statusToText
}
func (cust Customer) ToDto() dto.CustomerResponse {
	response := dto.CustomerResponse{
		Id:          cust.Id,
		Name:        cust.Name,
		City:        cust.City,
		ZipCode:     cust.ZipCode,
		DateOfBirth: cust.DateOfBirth,
		Status:      cust.statusToText(),
	}
	return response
}

type CustomerRepository interface {
	FindAll() ([]Customer, *errs.AppError)
	FindAllByStatus(string) ([]Customer, *errs.AppError)
	GetCustomer(string) (*Customer, *errs.AppError)
}
