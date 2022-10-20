package customer

import (
	"net/http"

	"github.com/Theofilush/warung-makan/delivery/middleware"
	_ "github.com/Theofilush/warung-makan/docs"
	"github.com/Theofilush/warung-makan/model"

	//ucc "github.com/Theofilush/warung-makan/usecase"
	useCaseCust "github.com/Theofilush/warung-makan/usecase/customer"
	"github.com/Theofilush/warung-makan/utils"
	"github.com/gin-gonic/gin"
	// swaggerFiles "github.com/swaggo/files"
	// ginSwagger "github.com/swaggo/gin-swagger"
)

type CustomerController struct {
	// router          *gin.Engine
	rgg             *gin.RouterGroup
	customerUsecase useCaseCust.CustomerUsecase
	// authUseCase     ucc.AuthUseCase
}

func (cc *CustomerController) userAuth(ctx *gin.Context) {
	var user model.UserCredential
	if err2 := ctx.ShouldBindJSON(&user); err2 != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"message": "can't bind struct",
		})
		return
	}

	token, err := cc.customerUsecase.UserAuth(user)

	if err != nil {
		ctx.AbortWithStatus(401)
		return
	}
	ctx.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}

func (cc *CustomerController) getAllCustomer(ctx *gin.Context) {
	customers, err := cc.customerUsecase.GetAllCustomer()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}
func (cc *CustomerController) getCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	customers, err := cc.customerUsecase.FindCustomerById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}

func (cc *CustomerController) registerCustomer(ctx *gin.Context) {
	var customer model.Customer
	customer.Id = utils.UuidGenerate()
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := cc.customerUsecase.RegisterCustomer(customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) UpdateCustomer(ctx *gin.Context) {
	var customer model.Customer
	if err := ctx.ShouldBindJSON(&customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := cc.customerUsecase.UpdateCustomer(customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	err := cc.customerUsecase.DeleteCustomer(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, err)
}

func NewCustomerController(routerGroup *gin.RouterGroup, usecaseCust useCaseCust.CustomerUsecase, tokenMdw middleware.AuthTokenMiddleware) *CustomerController {
	controllerr := CustomerController{
		rgg:             routerGroup,
		customerUsecase: usecaseCust,
	}

	// controllerr.rgg.POST("/auth", controllerr.userAuth)

	protectedGroup := controllerr.rgg.Group("/private", tokenMdw.RequireToken())

	protectedGroup.GET("/customer", controllerr.getAllCustomer)
	protectedGroup.GET("/customer/:id", controllerr.getCustomerById)
	protectedGroup.POST("/customer", controllerr.registerCustomer)
	protectedGroup.PUT("/customer", controllerr.UpdateCustomer)
	protectedGroup.DELETE("customer/:id", controllerr.DeleteCustomer)

	// ro := gin.Default()
	// controllerr.rgg.Any("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	// r.Run(":8080")
	return &controllerr
}
