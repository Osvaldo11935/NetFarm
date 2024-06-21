package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IOrderItemService interface {
	CreateOrderItem(OrderItem *entities.OrderItem) *common.ErrorResponse
	FindOrderItemByOrderId(orderId string, pagination *common.Pagination) ([]entities.OrderItem, *common.ErrorResponse)
	FindOrderItemById(orderItemId string) (entities.OrderItem, *common.ErrorResponse)
	UpdateOrderItem(orderItemId string, quantity *int) *common.ErrorResponse
	DeleteOrderItem(orderItemId string) *common.ErrorResponse
}
