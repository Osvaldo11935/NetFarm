package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IAddressService interface {
	CreateAddress(Address *entities.Address) *common.ErrorResponse
	FindAddressByUserId(userId string) (*entities.Address, *common.ErrorResponse)

	UpdateAddress(userId string, country *string, state *string, city *string,
		neighborhood *string, street *string, number *string, complement *string) *common.ErrorResponse

	DeleteAddress(userId string) *common.ErrorResponse
}
