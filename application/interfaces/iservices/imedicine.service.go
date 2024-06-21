package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IMedicineService interface {
	Create(medicine *entities.Medicine) *common.ErrorResponse
	FindAllMedicine(pagination *common.Pagination) ([]entities.Medicine, *common.ErrorResponse)
	FindMedicineByProvider(providerId string, pagination *common.Pagination) ([]entities.Medicine, *common.ErrorResponse)
	FindMedicineById(medicineId string) (entities.Medicine, *common.ErrorResponse)
	UpdateMedicine(medicineId string, name *string, description *string, quantity *int, price *float64) *common.ErrorResponse
	DeleteMedicine(fileTypeId string) *common.ErrorResponse
}
