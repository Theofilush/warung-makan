package delivery

import (
	"fmt"

	"github.com/Theofilush/warung-makan/config"
	"github.com/Theofilush/warung-makan/manager"
	"github.com/Theofilush/warung-makan/model"
	"github.com/Theofilush/warung-makan/utils/authenticator"
)

type Cli struct {
	useCaseManager manager.UseCaseManager
}

func (c *Cli) Run() {
	customerO2 := model.Customer{
		Id:            "6",
		Customer_name: "Doni",
		Address:       "Pasar Rebo",
	}

	customerUseCase := c.useCaseManager.CustomerUseCase()
	customerUseCase.RegisterCustomer(customerO2)
	allCustomer, _ := customerUseCase.GetAllCustomer()

	//repo := repository.NewCustomerRepository()
	//repoKampret := repository.NewCustomerKampretRepository()
	//repo.Create(customerO1)
	//repo.Create(customerO2)
	//
	//fmt.Println(repo.RetrieveAll())
	//customerUseCase := usecase.NewCustomerUseCase(repo)
	//customerUseCase.RegisterCustomer(customerO1)
	//customerUseCase.RegisterCustomer(customerO2)
	//cust := customerUseCase.FindCustomerById("2")
	//fmt.Println(cust)
	//allCustomer := customerUseCase.GetAllCustomer()
	fmt.Println(allCustomer)
}
func Console() *Cli {
	c := config.NewConfig()

	tokenService := authenticator.NewAccessToken(c.TokenConfig)
	// authUseCase := usecase.NewAuthUseCase(tokenService)

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepositoryManager(infra)
	usecase := manager.NewUseCaseManager(repo, tokenService)
	return &Cli{useCaseManager: usecase}
}
