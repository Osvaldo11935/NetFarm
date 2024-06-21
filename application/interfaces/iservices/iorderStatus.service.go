package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IOrderStatusService interface {
	CreateOrderStatus(OrderStatus *entities.OrderStatus) *common.ErrorResponse
	FindAllOrderStatus(pagination *common.Pagination) ([]entities.OrderStatus, *common.ErrorResponse)
	FindOrderStatusById(OrderStatusId string) (entities.OrderStatus, *common.ErrorResponse)
	UpdateOrderStatus(OrderStatusId string, name *string, description *string) *common.ErrorResponse
	DeleteOrderStatus(OrderStatusId string) *common.ErrorResponse
}
