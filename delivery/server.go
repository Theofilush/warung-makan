package delivery

import (
	"fmt"
	"net/http"

	"github.com/Theofilush/warung-makan/config"
	"github.com/Theofilush/warung-makan/delivery/controller/customer"
	"github.com/Theofilush/warung-makan/delivery/controller/menu"
	"github.com/Theofilush/warung-makan/delivery/middleware"
	"github.com/Theofilush/warung-makan/manager"
	"github.com/Theofilush/warung-makan/model"
	"github.com/Theofilush/warung-makan/usecase"
	"github.com/Theofilush/warung-makan/utils/authenticator"
	"github.com/gin-gonic/gin"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

type Server struct {
	useCaseManager manager.UseCaseManager
	engine         *gin.Engine
	host           string
	authUseCase    usecase.AuthUseCase
	tokenService   authenticator.AccessToken
}

func (s *Server) Run() {
	s.initController()
	err := s.engine.Run(s.host)
	if err != nil {
		panic(err)
	}
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

	publicRoute.Any("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	publicRoute.POST("/auth", func(c *gin.Context) {
		var user model.UserCredential
		if err := c.BindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"message": "can't bind struct",
			})
			return
		}
		if user.Username == "enigma" && user.Password == "123" {
			token, err := s.tokenService.CreateAccessToken(&user)
			if err != nil {
				c.AbortWithStatus(401)
			}
			c.JSON(200, gin.H{
				"token": token,
			})
		} else {
			c.AbortWithStatus(401)
		}

	})

	// controller.NewAppController(publicRoute, s.authUseCase, tokenMdw)

	customer.NewCustomerController(publicRoute, s.useCaseManager.CustomerUseCase(), tokenMdw)
	menu.NewMenuController(publicRoute, s.useCaseManager.MenuUseCase(), tokenMdw)
}
