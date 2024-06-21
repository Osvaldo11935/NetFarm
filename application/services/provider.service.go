package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type ProviderService struct {
	ProviderRepo irepositories.IProviderRepository
}

func NewProviderService(ProviderRepo irepositories.IProviderRepository) iservices.IProviderService {
	return &ProviderService{
		ProviderRepo: ProviderRepo,
	}
}

func (s *ProviderService) Create(provider *entities.Provider) *common.ErrorResponse {

	createdErr := s.ProviderRepo.Create(provider)

	if createdErr != nil {
		return errors.NewProviderUnknownError(createdErr.Error(), "falha ao cadastrar fornecedor")
	}

	return nil
}

func (s *ProviderService) FindAllProvider(pagination *common.Pagination) ([]entities.Provider, *common.ErrorResponse) {

	var providers []entities.Provider

	query := s.ProviderRepo.Query()

	query = query.Scopes(extensions.Paginate(&providers, pagination, query))

	err := query.Find(&providers).Error

	if err != nil {
		return providers, errors.NewProviderUnknownError(err.Error(), "falha ao buscar fornecedor")
	}

	pagination.Data = providers

	return providers, nil
}

func (s *ProviderService) FindProviderById(providerId string) (entities.Provider, *common.ErrorResponse) {

	var provider entities.Provider

	err := s.ProviderRepo.Query().First(&provider, "Id", providerId).Error

	if err != nil {

		return provider, errors.NewProviderUnknownError(err.Error(), "falha ao buscar fornecedor")
	}

	return provider, nil
}

func (s *ProviderService) UpdateProvider(providerId string, name *string) *common.ErrorResponse {

	provider, err := s.FindProviderById(providerId)

	if err != nil {
		return err
	}

	provider.Update(name)

	updatedErr := s.ProviderRepo.Update(provider)

	if updatedErr != nil {
		return errors.NewProviderUnknownError(updatedErr.Error(), "falha ao atualizar fornecedor")
	}

	return nil
}

func (s *ProviderService) DeleteProvider(providerId string) *common.ErrorResponse {

	provider, err := s.FindProviderById(providerId)

	if err != nil {
		return err
	}

	deletedErr := s.ProviderRepo.Delete(provider)

	if deletedErr != nil {
		return errors.NewProviderUnknownError(deletedErr.Error(), "falha ao deletar fornecedor")
	}

	return nil
}
