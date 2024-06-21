package repositories

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories/common"
)

type FileTypeRepository struct {
	*common.Repository
}

func NewFileTypeRepository() irepositories.IFileTypeRepository {
	return &FileTypeRepository{
		Repository: common.NewRepository().(*common.Repository),
	}
}
