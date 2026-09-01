package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type MedicineRepository struct {
	*common.Repository
}

func NewMedicineRepository(db *gorm.DB) irepositories.IMedicineRepository {
	return &MedicineRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}
