package manager

import "enigmacamp.com/final-project/usecase/customer"

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
