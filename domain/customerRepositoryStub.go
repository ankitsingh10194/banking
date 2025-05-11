package domain

type CustomerRepositoryStub struct {
	customer []Customer
}

func (c CustomerRepositoryStub) FindAll() ([]Customer, error) {
	return c.customer, nil
}

func NewCustomerRepositoryStub() CustomerRepositoryStub {
	customers := []Customer{
		{Id: "1", Name: "Ankit", City: "Jalaun", ZipCode: "285123", DateOfBirth: "10-01-1994", Status: "1"},
		{Id: "2", Name: "Sonu", City: "Jalaun", ZipCode: "285123", DateOfBirth: "10-01-1994", Status: "1"},
	}
	return CustomerRepositoryStub{customer: customers}
}
