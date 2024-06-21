package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
	"NetFarm/shared/models/responses"
)

type IOrderDetailService interface {
	Calculate(items []entities.OrderItem) responses.OrderDetailResponse
	CreateOrderDetail(orderDetail *entities.OrderDetail) *common.ErrorResponse
	FindAllOrderDetail(pagination *common.Pagination) ([]entities.OrderDetail, *common.ErrorResponse)
}
