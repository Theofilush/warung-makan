package manager

import (
	"github.com/Theofilush/warung-makan/repository/customer"
	"github.com/Theofilush/warung-makan/repository/menu"
	"github.com/Theofilush/warung-makan/repository/order"
	ord "github.com/Theofilush/warung-makan/repository/order"
)

type RepositoryManager interface {
	CustomerRepo() customer.CustomerRepository
	MenuRepo() menu.MenuRepository
	OrderRepo() order.OrderRepository
}

type repositoryManager struct {
	infra InfraManager
}

func (r *repositoryManager) CustomerRepo() customer.CustomerRepository {
	return customer.NewCustomerDbRepository(r.infra.DbConn())
}

func (r *repositoryManager) MenuRepo() menu.MenuRepository {
	return menu.NewMenuDbRepository(r.infra.DbConn())
}
func (r *repositoryManager) OrderRepo() ord.OrderRepository {
	return ord.NewOrderDbRepository(r.infra.DbConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repositoryManager{infra: manager}
}
