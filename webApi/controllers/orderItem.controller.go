package controllers

import (
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	common2 "NetFarm/shared/models/common"
	"NetFarm/webApi/models/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderItemController struct {
	OrderItemService iservices.IOrderItemService
}

func NewOrderItemController(service iservices.IOrderItemService) *OrderItemController {
	return &OrderItemController{
		OrderItemService: service,
	}
}

func (controller *OrderItemController) HandleCreate(ctx *gin.Context) {

	createOrderItemRequest := entities.OrderItem{}
	err := ctx.ShouldBindJSON(&createOrderItemRequest)

	if err != nil {
		fmt.Println(err)
	}

	orderId := ctx.Param("orderId")

	createOrderItemRequest.OrderId = orderId

	createErr := controller.OrderItemService.CreateOrderItem(&createOrderItemRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common2.BaseCreateResponse{Id: createOrderItemRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderItemController) HandleFindOrderItemByOrderId(ctx *gin.Context) {

	orderId := ctx.Param("orderId")

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common2.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.OrderItemService.FindOrderItemByOrderId(orderId, &pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *OrderItemController) HandleFindOrderItemById(ctx *gin.Context) {
	OrderItemId := ctx.Param("OrderItemId")

	OrderItem, findErr := controller.OrderItemService.FindOrderItemById(OrderItemId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, OrderItem)
}

func (controller *OrderItemController) HandleUpdateOrderItem(ctx *gin.Context) {

	orderItemId := ctx.Param("orderItemId")

	createOrderItemRequest := requests.UpdateOrderItemRequest{}

	err := ctx.ShouldBindJSON(&createOrderItemRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.OrderItemService.UpdateOrderItem(
		orderItemId,
		createOrderItemRequest.Quantity)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common2.BaseCreateResponse{Id: orderItemId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderItemController) HandleDeleteOrderItem(ctx *gin.Context) {

	orderItemId := ctx.Param("orderItemId")

	deletedErr := controller.OrderItemService.DeleteOrderItem(orderItemId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
