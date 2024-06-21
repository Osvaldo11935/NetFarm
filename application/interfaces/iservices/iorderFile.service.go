package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IOrderFileService interface {
	CreateOrderFile(OrderFile *entities.OrderFile) *common.ErrorResponse
	FindOrderFileByOrderId(orderId string, pagination *common.Pagination) ([]entities.OrderFile, *common.ErrorResponse)
	FindOrderFileById(orderFileId string) (entities.OrderFile, *common.ErrorResponse)
	DeleteOrderFile(orderFileId string) *common.ErrorResponse
}
