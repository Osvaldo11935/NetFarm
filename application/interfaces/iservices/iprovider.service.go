package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IProviderService interface {
	Create(provider *entities.Provider) *common.ErrorResponse
	FindAllProvider(pagination *common.Pagination) ([]entities.Provider, *common.ErrorResponse)
	FindProviderById(providerId string) (entities.Provider, *common.ErrorResponse)
	UpdateProvider(providerId string, name *string) *common.ErrorResponse
	DeleteProvider(providerId string) *common.ErrorResponse
}
