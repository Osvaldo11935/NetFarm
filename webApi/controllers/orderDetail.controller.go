package controllers

import (
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderDetailController struct {
	OrderDetailService iservices.IOrderDetailService
}

func NewOrderDetailController(service iservices.IOrderDetailService) *OrderDetailController {
	return &OrderDetailController{
		OrderDetailService: service,
	}
}

func (controller *OrderDetailController) HandleCreate(ctx *gin.Context) {

	createOrderDetailRequest := entities.OrderDetail{}
	err := ctx.ShouldBindJSON(&createOrderDetailRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.OrderDetailService.CreateOrderDetail(&createOrderDetailRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createOrderDetailRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderDetailController) HandleOrderCalculate(ctx *gin.Context) {

	var createOrderItemRequest []entities.OrderItem
	err := ctx.ShouldBindJSON(&createOrderItemRequest)

	if err != nil {
		fmt.Println(err)
	}

	calculateResult := controller.OrderDetailService.Calculate(createOrderItemRequest)

	//if createErr != nil {
	//	ctx.JSON(http.StatusBadRequest, createErr)
	//	return
	//}

	ctx.JSON(http.StatusOK, calculateResult)
}

func (controller *OrderDetailController) HandleFindAllOrderDetails(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.OrderDetailService.FindAllOrderDetail(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}
