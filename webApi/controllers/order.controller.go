package controllers

import (
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
	"NetFarm/webApi/models/responses"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderController struct {
	orderService iservices.IOrderService
}

func NewOrderController(service iservices.IOrderService) *OrderController {
	return &OrderController{
		orderService: service,
	}
}

func (controller *OrderController) HandleCreate(ctx *gin.Context) {

	createOrderRequest := entities.Order{}
	err := ctx.ShouldBindJSON(&createOrderRequest)

	userId := ctx.Param("userId")

	createOrderRequest.UserId = userId

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.orderService.CreateOrder(&createOrderRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := responses.CreateOrderResponse{
		Id: createOrderRequest.Id,
	}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderController) HandleFindAllOrders(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.orderService.FindAllOrder(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *OrderController) HandleFindOrdersByUserId(ctx *gin.Context) {

	userId := ctx.Param("userId")

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	orders, findErr := controller.orderService.FindOrderByUserId(userId, &pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, orders)
}

func (controller *OrderController) HandleFindOrderById(ctx *gin.Context) {
	orderId := ctx.Param("orderId")

	Order, findErr := controller.orderService.FindOrderById(orderId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, Order)
}

func (controller *OrderController) HandleUpdateStatusOrder(ctx *gin.Context) {

	orderId := ctx.Param("orderId")
	statusId := ctx.Param("statusId")

	updatedErr := controller.orderService.UpdateOrder(orderId, &statusId)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: orderId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderController) HandleDeleteOrder(ctx *gin.Context) {

	orderId := ctx.Param("orderId")

	deletedErr := controller.orderService.DeleteOrder(orderId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
