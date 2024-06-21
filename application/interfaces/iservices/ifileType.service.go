package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IFileTypeService interface {
	CreateFileType(FileType *entities.FileType) *common.ErrorResponse
	FindAllFileType(pagination *common.Pagination) ([]entities.FileType, *common.ErrorResponse)
	FindFileTypeById(fileTypeId string) (entities.FileType, *common.ErrorResponse)
	UpdateFileType(fileTypeId string, name *string, description *string) *common.ErrorResponse
	DeleteFileType(fileTypeId string) *common.ErrorResponse
}
