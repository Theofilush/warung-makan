package customer

import (
	"enigmacamp.com/final-project/model"
	"enigmacamp.com/final-project/repository/customer"
)

type CustomerUsecase interface {
	RegisterCustomer(customer model.Customer) error
	FindCustomerById(id string) (model.Customer, error)
	GetAllCustomer() ([]model.Customer, error)
	UpdateCustomer(customer model.Customer) error
	DeleteCustomer(id string) error
}
type customerUsecase struct {
	repo customer.CustomerRepository
}

func (c *customerUsecase) RegisterCustomer(customer model.Customer) error {
	return c.repo.Create(customer)
}

func (c *customerUsecase) FindCustomerById(id string) (model.Customer, error) {
	return c.repo.FindById(id)
}

func (c *customerUsecase) GetAllCustomer() ([]model.Customer, error) {
	return c.repo.RetrieveAll()
}

func (c *customerUsecase) UpdateCustomer(customer model.Customer) error {
	return c.repo.Update(customer)
}

func (c *customerUsecase) DeleteCustomer(id string) error {
	return c.repo.Delete(id)
}

func NewCustomerUseCase(repo customer.CustomerRepository) CustomerUsecase {
	return &customerUsecase{repo: repo}
}
