package service

import (
	"github.com/ankitsingh10194/banking/domain"
	"github.com/ankitsingh10194/banking/dto"
	"github.com/ankitsingh10194/banking/errs"
)

type CustomerService interface {
	GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError)
	GetAllCustomersByStatus(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(customerId string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerservice struct {
	repo domain.CustomerRepository
}

func (c DefaultCustomerservice) GetAllCustomers() ([]dto.CustomerResponse, *errs.AppError) {
	cust, err := c.repo.FindAll()
	if err != nil {
		return nil, err
	}
	response := []dto.CustomerResponse{}
	for _, c := range cust {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (c DefaultCustomerservice) GetAllCustomersByStatus(status string) ([]dto.CustomerResponse, *errs.AppError) {
	cust, err := c.repo.FindAllByStatus(status)
	if err != nil {
		return nil, err
	}
	response := make([]dto.CustomerResponse, 0)
	for _, c := range cust {
		response = append(response, c.ToDto())
	}
	return response, nil
}

func (c DefaultCustomerservice) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	cust, err := c.repo.GetCustomer(id)
	if err != nil {
		return nil, err
	}
	response := cust.ToDto()
	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepositoryDb) DefaultCustomerservice {
	return DefaultCustomerservice{repo: repository}
}
