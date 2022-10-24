package order

import (
	"github.com/Theofilush/warung-makan/model"
	ord "github.com/Theofilush/warung-makan/repository/order"
	"github.com/Theofilush/warung-makan/utils/authenticator"
)

type OrderUsecase interface {
	RegisterOrder(order model.Order) error
	FindOrderById(id string) (model.Order_details2, error)
	GetAllOrder() ([]model.Order_details2, error)
	UpdateOrder(order model.Order) error
	DeleteOrder(id string) error
}
type orderUsecase struct {
	repo         ord.OrderRepository
	tokenService authenticator.AccessToken
}

func (c *orderUsecase) RegisterOrder(order model.Order) error {
	return c.repo.Create(order)
}

func (c *orderUsecase) FindOrderById(id string) (model.Order_details2, error) {
	return c.repo.FindById(id)
}

func (c *orderUsecase) GetAllOrder() ([]model.Order_details2, error) {
	return c.repo.RetrieveAll()
}

func (c *orderUsecase) UpdateOrder(order model.Order) error {
	return c.repo.Update(order)
}

func (c *orderUsecase) DeleteOrder(id string) error {
	return c.repo.Delete(id)
}

func NewOrderUseCase(repo ord.OrderRepository, service authenticator.AccessToken) OrderUsecase {
	orderUsecase := new(orderUsecase)
	orderUsecase.tokenService = service
	orderUsecase.repo = repo

	return orderUsecase

	// return &orderUsecase{repo: repo, tokenService: service}
}
