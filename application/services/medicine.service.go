package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type MedicineService struct {
	MedicineRepo irepositories.IMedicineRepository
}

func NewMedicineService(MedicineRepo irepositories.IMedicineRepository) iservices.IMedicineService {
	return &MedicineService{
		MedicineRepo: MedicineRepo,
	}
}

func (s *MedicineService) Create(medicine *entities.Medicine) *common.ErrorResponse {

	createdErr := s.MedicineRepo.Create(medicine)

	if createdErr != nil {
		return errors.NewMedicineUnknownError(createdErr.Error(), "falha ao cadastrar remedio")
	}

	return nil
}

func (s *MedicineService) FindAllMedicine(pagination *common.Pagination) ([]entities.Medicine, *common.ErrorResponse) {

	query := s.MedicineRepo.Query()

	var medicines []entities.Medicine

	query = query.Scopes(extensions.Paginate(&medicines, pagination, query))

	findErr := query.Find(&medicines).Error

	if findErr != nil {
		return medicines, errors.NewMedicineUnknownError(findErr.Error(), "falha ao buscar remedio")
	}
	
	pagination.Data = medicines

	return medicines, nil
}

func (s *MedicineService) FindMedicineByProvider(providerId string, pagination *common.Pagination) ([]entities.Medicine, *common.ErrorResponse) {

	query := s.MedicineRepo.Query().Where("ProviderId", providerId)

	var medicines []entities.Medicine

	query = query.Scopes(extensions.Paginate(&medicines, pagination, query))

	findErr := query.Find(&medicines).Error

	if findErr != nil {
		return medicines, errors.NewMedicineUnknownError(findErr.Error(), "falha ao buscar remedio")
	}

	return medicines, nil
}

func (s *MedicineService) FindMedicineById(medicineId string) (entities.Medicine, *common.ErrorResponse) {

	var medicine entities.Medicine

	findErr := s.MedicineRepo.Query().First(&medicine, "Id", medicineId).Error

	if findErr != nil {
		return medicine, errors.NewMedicineUnknownError(findErr.Error(), "falha ao buscar remedio")
	}

	return medicine, nil
}

func (s *MedicineService) UpdateMedicine(medicineId string, name *string, description *string, quantity *int, price *float64) *common.ErrorResponse {

	medicine, err := s.FindMedicineById(medicineId)

	if err != nil {
		return err
	}

	medicine.Update(name, description, quantity, price)

	updatedErr := s.MedicineRepo.Update(medicine)

	if updatedErr != nil {
		return errors.NewMedicineUnknownError(updatedErr.Error(), "falha ao atualizar remedio")
	}

	return nil
}

func (s *MedicineService) DeleteMedicine(fileTypeId string) *common.ErrorResponse {

	medicine, err := s.FindMedicineById(fileTypeId)

	if err != nil {
		return err
	}

	deletedErr := s.MedicineRepo.Delete(medicine)

	if deletedErr != nil {
		return errors.NewMedicineUnknownError(deletedErr.Error(), "falha ao deletar remedio")
	}

	return nil
}
