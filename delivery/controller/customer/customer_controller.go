package customer

import (
	"net/http"

	"enigmacamp.com/final-project/model"
	"enigmacamp.com/final-project/usecase/customer"
	"enigmacamp.com/final-project/utils"
	"github.com/gin-gonic/gin"
)

type CustomerController struct {
	router  *gin.Engine
	usecase customer.CustomerUsecase
}

func (cc *CustomerController) getAllCustomer(ctx *gin.Context) {
	customers, err := cc.usecase.GetAllCustomer()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customers)
}
func (cc *CustomerController) getCustomerById(ctx *gin.Context) {
	id := ctx.Param("id")
	customers, err := cc.usecase.FindCustomerById(id)
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
	if err := cc.usecase.RegisterCustomer(customer); err != nil {
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
	if err := cc.usecase.UpdateCustomer(customer); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, customer)
}

func (cc *CustomerController) DeleteCustomer(ctx *gin.Context) {
	id := ctx.Param("id")
	err := cc.usecase.DeleteCustomer(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, err)
}

func NewCustomerController(r *gin.Engine, usecase customer.CustomerUsecase) *CustomerController {
	controller := CustomerController{
		router:  r,
		usecase: usecase,
	}
	r.GET("/customer", controller.getAllCustomer)
	r.GET("/customer/:id", controller.getCustomerById)
	r.POST("/customer", controller.registerCustomer)
	r.PUT("/customer", controller.UpdateCustomer)
	r.DELETE("customer/:id", controller.DeleteCustomer)
	return &controller
}
