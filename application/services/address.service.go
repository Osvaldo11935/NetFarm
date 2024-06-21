package services

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type AddressService struct {
	AddressRepo irepositories.IAddressRepository
}

func NewAddressService(AddressRepo irepositories.IAddressRepository) iservices.IAddressService {
	return &AddressService{
		AddressRepo: AddressRepo,
	}
}

func (s *AddressService) CreateAddress(Address *entities.Address) *common.ErrorResponse {
	address, findAddressErr := s.FindAddressByUserId(Address.UserId)

	if findAddressErr != nil {
		return findAddressErr
	}

	if address != nil {
		return errors.NewAddressAlreadyExists("falha ao cadastrar endereço do usuario")
	}

	err := s.AddressRepo.Create(Address)

	if err != nil {
		return errors.NewAddressUnknownError(err.Error(), "falha ao cadastrar endereço do usuario")
	}

	return nil
}

func (s *AddressService) FindAddressByUserId(userId string) (*entities.Address, *common.ErrorResponse) {
	var address entities.Address

	err := s.AddressRepo.Query().First(&address, "UserId", userId).Error

	if err != nil {
		return nil, errors.NewAddressUnknownError(err.Error(), "falha ao buscar endereço do usuario")
	}
	return &address, nil
}

func (s *AddressService) UpdateAddress(userId string, country *string, state *string, city *string,
	neighborhood *string, street *string, number *string, complement *string) *common.ErrorResponse {

	address, err := s.FindAddressByUserId(userId)

	if err != nil {
		return err
	}

	address.Update(country, state, city, neighborhood, street, number, complement)

	updatedErr := s.AddressRepo.Update(address)

	if updatedErr != nil {
		return errors.NewAddressUnknownError(updatedErr.Error(), "falha ao atualizar dados do endereço do usuario")
	}

	return nil
}

func (s *AddressService) DeleteAddress(userId string) *common.ErrorResponse {

	Person, findAddressErr := s.FindAddressByUserId(userId)

	if findAddressErr != nil {
		return findAddressErr
	}

	deleteErr := s.AddressRepo.Delete(Person)

	if deleteErr != nil {
		return errors.NewAddressUnknownError(deleteErr.Error(), "falha ao deletar endereço do usuario")
	}

	return nil
}
