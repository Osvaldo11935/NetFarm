package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type ImposedService struct {
	ImposedRepo irepositories.IImposedRepository
}

func NewImposedService(ImposedRepo irepositories.IImposedRepository) iservices.IImposedService {
	return &ImposedService{
		ImposedRepo: ImposedRepo,
	}
}

func (s *ImposedService) CreateImposed(imposed *entities.Imposed) *common.ErrorResponse {

	createdErr := s.ImposedRepo.Create(imposed)

	if createdErr != nil {
		return errors.NewImposedUnknownError(createdErr.Error(), "falha ao cadastrar imposto")
	}

	return nil
}

func (s *ImposedService) FindAllImposed(pagination *common.Pagination) ([]entities.Imposed, *common.ErrorResponse) {

	var Imposeds []entities.Imposed

	query := s.ImposedRepo.Query()

	query = query.Scopes(extensions.Paginate(&Imposeds, pagination, query))

	findErr := query.Find(&Imposeds).Error

	if findErr != nil {
		return Imposeds, errors.NewImposedUnknownError(findErr.Error(), "falha ao buscar imposto")
	}

	pagination.Data = Imposeds

	return Imposeds, nil
}

func (s *ImposedService) FindImposedById(imposedId string) (entities.Imposed, *common.ErrorResponse) {

	var Imposed entities.Imposed

	findErr := s.ImposedRepo.Query().First(&Imposed, "Id", imposedId).Error

	if findErr != nil {
		return Imposed, errors.NewImposedUnknownError(findErr.Error(), "falha ao buscar imposto")
	}

	return Imposed, nil
}

func (s *ImposedService) UpdateImposed(imposedId string, name *string, rate *float64) *common.ErrorResponse {

	Imposed, err := s.FindImposedById(imposedId)

	if err != nil {
		return err
	}

	Imposed.Update(name, rate)

	updatedErr := s.ImposedRepo.Update(Imposed)

	if updatedErr != nil {
		return errors.NewImposedUnknownError(updatedErr.Error(), "falha ao atualizar imposto")
	}

	return nil
}

func (s *ImposedService) DeleteImposed(imposedId string) *common.ErrorResponse {

	Imposed, err := s.FindImposedById(imposedId)

	if err != nil {
		return err
	}

	deletedErr := s.ImposedRepo.Delete(Imposed)

	if deletedErr != nil {
		return errors.NewImposedUnknownError(deletedErr.Error(), "falha ao deletar imposto")
	}

	return nil
}
