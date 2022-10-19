package customer

import (
	"github.com/Theofilush/warung-makan/model"
	"github.com/Theofilush/warung-makan/repository/customer"
	"github.com/Theofilush/warung-makan/utils/authenticator"
)

type CustomerUsecase interface {
	UserAuth(user model.UserCredential) (token string, err error)
	RegisterCustomer(customer model.Customer) error
	FindCustomerById(id string) (model.Customer, error)
	GetAllCustomer() ([]model.Customer, error)
	UpdateCustomer(customer model.Customer) error
	DeleteCustomer(id string) error
}
type customerUsecase struct {
	repo         customer.CustomerRepository
	tokenService authenticator.AccessToken
}

// type authUseCase struct {
// 	tokenService authenticator.AccessToken
// }

func (c *customerUsecase) UserAuth(user model.UserCredential) (token string, err error) {
	if user.Username == "enigma" && user.Password == "123" {
		token, err := c.tokenService.CreateAccessToken(&user)
		if err != nil {
			return "", err
		}
		return token, nil
	} else {
		return "", nil
	}
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

func NewCustomerUseCase(repo customer.CustomerRepository, service authenticator.AccessToken) CustomerUsecase {
	customerUsecase := new(customerUsecase)
	customerUsecase.tokenService = service
	customerUsecase.repo = repo

	return customerUsecase

	// return &customerUsecase{repo: repo, tokenService: service}
}
