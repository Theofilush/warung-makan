package manager

import (
	"github.com/Theofilush/warung-makan/repository/customer"
	"github.com/Theofilush/warung-makan/repository/menu"
)

type RepositoryManager interface {
	CustomerRepo() customer.CustomerRepository
	MenuRepo() menu.MenuRepository
}

type repoistoryManager struct {
	infra InfraManager
}

func (r *repoistoryManager) CustomerRepo() customer.CustomerRepository {
	return customer.NewCustomerDbRepository(r.infra.DbConn())
}

func (r *repoistoryManager) MenuRepo() menu.MenuRepository {
	return menu.NewMenuDbRepository(r.infra.DbConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repoistoryManager{infra: manager}
}
