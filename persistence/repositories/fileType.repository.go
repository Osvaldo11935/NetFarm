package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
	"gorm.io/gorm"
)

type FileTypeRepository struct {
	*common.Repository
}

func NewFileTypeRepository(db *gorm.DB) irepositories.IFileTypeRepository {
	return &FileTypeRepository{
		Repository: common.NewRepository(db).(*common.Repository),
	}
}
