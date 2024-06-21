package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IImposedService interface {
	CreateImposed(imposed *entities.Imposed) *common.ErrorResponse
	FindAllImposed(pagination *common.Pagination) ([]entities.Imposed, *common.ErrorResponse)
	FindImposedById(imposedId string) (entities.Imposed, *common.ErrorResponse)
	UpdateImposed(imposedId string, name *string, rate *float64) *common.ErrorResponse
	DeleteImposed(imposedId string) *common.ErrorResponse
}
