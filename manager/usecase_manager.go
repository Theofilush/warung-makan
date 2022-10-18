package manager

import "github.com/Theofilush/warung-makan/usecase/customer"

type UseCaseManager interface {
	CustomerUseCase() customer.CustomerUsecase
}

type useCaseManager struct {
	repoManager RepositoryManager
}

func (u *useCaseManager) CustomerUseCase() customer.CustomerUsecase {
	return customer.NewCustomerUseCase(u.repoManager.CustomerRepo())
}

func NewUseCaseManager(repoManager RepositoryManager) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
	}
}
