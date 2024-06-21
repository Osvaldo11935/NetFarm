package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type MedicineRepository struct {
	*common.Repository
}

func NewMedicineRepository() irepositories.IMedicineRepository {
	return &MedicineRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}
