package manager

import (
	"github.com/Theofilush/warung-makan/usecase"
	"github.com/Theofilush/warung-makan/usecase/customer"
	"github.com/Theofilush/warung-makan/utils/authenticator"
)

type UseCaseManager interface {
	CustomerUseCase() customer.CustomerUsecase
	AuthUseCase() usecase.AuthUseCase
}

type useCaseManager struct {
	repoManager RepositoryManager
	service     authenticator.AccessToken
}

func (u *useCaseManager) CustomerUseCase() customer.CustomerUsecase {
	return customer.NewCustomerUseCase(u.repoManager.CustomerRepo(), u.service)
}

func (u *useCaseManager) AuthUseCase() usecase.AuthUseCase {
	return usecase.NewAuthUseCase(u.service)
}

func NewUseCaseManager(repoManager RepositoryManager, service authenticator.AccessToken) UseCaseManager {
	return &useCaseManager{
		repoManager: repoManager,
		service:     service,
	}
}
