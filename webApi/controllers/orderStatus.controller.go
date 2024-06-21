package controllers

import (
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
	"NetFarm/webApi/models/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderStatusController struct {
	OrderStatusService iservices.IOrderStatusService
}

func NewOrderStatusController(service iservices.IOrderStatusService) *OrderStatusController {
	return &OrderStatusController{
		OrderStatusService: service,
	}
}

func (controller *OrderStatusController) HandleCreate(ctx *gin.Context) {

	createOrderStatusRequest := entities.OrderStatus{}
	err := ctx.ShouldBindJSON(&createOrderStatusRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.OrderStatusService.CreateOrderStatus(&createOrderStatusRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createOrderStatusRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderStatusController) HandleFindAllOrderStatus(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.OrderStatusService.FindAllOrderStatus(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *OrderStatusController) HandleFindOrderStatusById(ctx *gin.Context) {
	OrderStatusId := ctx.Param("OrderStatusId")

	OrderStatus, findErr := controller.OrderStatusService.FindOrderStatusById(OrderStatusId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, OrderStatus)
}

func (controller *OrderStatusController) HandleUpdateOrderStatus(ctx *gin.Context) {

	orderStatusId := ctx.Param("orderStatusId")

	createOrderStatusRequest := requests.UpdateOrderStatusRequest{}

	err := ctx.ShouldBindJSON(&createOrderStatusRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.OrderStatusService.UpdateOrderStatus(
		orderStatusId,
		createOrderStatusRequest.Type,
		createOrderStatusRequest.Description)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: orderStatusId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderStatusController) HandleDeleteOrderStatus(ctx *gin.Context) {

	orderStatusId := ctx.Param("orderStatusId")

	deletedErr := controller.OrderStatusService.DeleteOrderStatus(orderStatusId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
