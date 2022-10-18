package manager

import "github.com/Theofilush/warung-makan/repository/customer"

type RepositoryManager interface {
	CustomerRepo() customer.CustomerRepository
}

type repoistoryManager struct {
	infra InfraManager
}

func (r *repoistoryManager) CustomerRepo() customer.CustomerRepository {
	return customer.NewCustomerDbRepository(r.infra.DbConn())
}

func NewRepositoryManager(manager InfraManager) RepositoryManager {
	return &repoistoryManager{infra: manager}
}
