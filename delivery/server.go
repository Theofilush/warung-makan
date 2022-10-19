package delivery

import (
	"fmt"

	"github.com/Theofilush/warung-makan/config"
	"github.com/Theofilush/warung-makan/delivery/controller/customer"
	"github.com/Theofilush/warung-makan/delivery/middleware"
	"github.com/Theofilush/warung-makan/manager"
	"github.com/Theofilush/warung-makan/usecase"
	"github.com/Theofilush/warung-makan/utils/authenticator"
	"github.com/gin-gonic/gin"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	authUseCase    usecase.AuthUseCase
	tokenService   authenticator.AccessToken
}

func NewServer() *Server {
	c := config.NewConfig()
	r := gin.Default()

	tokenService := authenticator.NewAccessToken(c.TokenConfig)
	authUseCase := usecase.NewAuthUseCase(tokenService)

	infra := manager.NewInfraManager(c)
	repo := manager.NewRepositoryManager(infra)
	usecasee := manager.NewUseCaseManager(repo, tokenService)

	if c.ApiHost == "" || c.ApiPort == "" {
		panic("No Host or port define")
	}
	host := fmt.Sprintf("%s:%s", c.ApiHost, c.ApiPort)
	return &Server{useCaseManager: usecasee, engine: r, host: host, authUseCase: authUseCase, tokenService: tokenService}
}

func (s *Server) initController() {
	publicRoute := s.engine.Group("/v1")
	tokenMdw := middleware.NewTokenValidator(s.tokenService)

	// controller.NewAppController(publicRoute, s.authUseCase, tokenMdw)

	customer.NewCustomerController(publicRoute, s.useCaseManager.CustomerUseCase(), tokenMdw)
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
}
