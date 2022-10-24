package order

import (
	"net/http"

	"github.com/Theofilush/warung-makan/delivery/middleware"
	_ "github.com/Theofilush/warung-makan/docs"
	"github.com/Theofilush/warung-makan/model"
	"github.com/Theofilush/warung-makan/utils"

	useCaseOrd "github.com/Theofilush/warung-makan/usecase/order"
	"github.com/gin-gonic/gin"
)

type OrderController struct {
	rgg          *gin.RouterGroup
	orderUsecase useCaseOrd.OrderUsecase
}

func (cc *OrderController) getAllOrder(ctx *gin.Context) {
	orders, err := cc.orderUsecase.GetAllOrder()
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (cc *OrderController) getOrderById(ctx *gin.Context) {
	id := ctx.Param("id")
	orders, err := cc.orderUsecase.FindOrderById(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, orders)
}

func (cc *OrderController) registerOrder(ctx *gin.Context) {
	var order model.Order
	order.Id = utils.UuidGenerate()
	if err := ctx.BindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := cc.orderUsecase.RegisterOrder(order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func (cc *OrderController) UpdateOrder(ctx *gin.Context) {
	var order model.Order
	if err := ctx.ShouldBindJSON(&order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	if err := cc.orderUsecase.UpdateOrder(order); err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, order)
}

func (cc *OrderController) DeleteOrder(ctx *gin.Context) {
	id := ctx.Param("id")
	err := cc.orderUsecase.DeleteOrder(id)
	if err != nil {
		ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"err": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, err)
}

func NewOrderController(routerGroup *gin.RouterGroup, usecaseOrd useCaseOrd.OrderUsecase, tokenMdw middleware.AuthTokenMiddleware) *OrderController {
	controllerr := OrderController{
		rgg:          routerGroup,
		orderUsecase: usecaseOrd,
	}

	protectedGroup := controllerr.rgg.Group("/private", tokenMdw.RequireToken())

	protectedGroup.GET("/order", controllerr.getAllOrder)      //
	protectedGroup.GET("/order/:id", controllerr.getOrderById) //
	protectedGroup.POST("/order", controllerr.registerOrder)   //
	// protectedGroup.PUT("/order", controllerr.UpdateOrder) // no fix coz no use update
	protectedGroup.DELETE("order/:id", controllerr.DeleteOrder)

	return &controllerr
}
