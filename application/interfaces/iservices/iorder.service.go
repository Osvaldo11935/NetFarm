package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IOrderService interface {
	CreateOrder(order *entities.Order) *common.ErrorResponse
	FindAllOrder(pagination *common.Pagination) ([]entities.Order, *common.ErrorResponse)
	FindOrderByUserId(userId string, pagination *common.Pagination) ([]entities.Order, *common.ErrorResponse)
	FindOrderById(orderId string) (entities.Order, *common.ErrorResponse)
	UpdateOrder(orderId string, statusId *string) *common.ErrorResponse
	DeleteOrder(OrderId string) *common.ErrorResponse
}
