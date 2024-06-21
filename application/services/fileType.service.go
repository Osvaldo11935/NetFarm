package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type FileTypeService struct {
	FileTypeRepo irepositories.IFileTypeRepository
}

func NewFileTypeService(FileTypeRepo irepositories.IFileTypeRepository) iservices.IFileTypeService {
	return &FileTypeService{
		FileTypeRepo: FileTypeRepo,
	}
}

func (s *FileTypeService) CreateFileType(FileType *entities.FileType) *common.ErrorResponse {

	createdErr := s.FileTypeRepo.Create(FileType)

	if createdErr != nil {
		return errors.NewFileTypeUnknownError(createdErr.Error(), "falha ao cadastrar tipo de ficheiro")
	}

	return nil
}

func (s *FileTypeService) FindAllFileType(pagination *common.Pagination) ([]entities.FileType, *common.ErrorResponse) {

	var fileTypes []entities.FileType

	query := s.FileTypeRepo.Query()

	query = query.Scopes(extensions.Paginate(&fileTypes, pagination, query))

	findErr := query.Find(&fileTypes).Error

	if findErr != nil {
		return fileTypes, errors.NewFileTypeUnknownError(findErr.Error(), "falha ao buscar tipo de ficheiro")
	}

	pagination.Data = fileTypes

	return fileTypes, nil
}

func (s *FileTypeService) FindFileTypeById(fileTypeId string) (entities.FileType, *common.ErrorResponse) {

	var fileType entities.FileType

	findErr := s.FileTypeRepo.Query().First(&fileType, "Id", fileTypeId).Error

	if findErr != nil {
		return fileType, errors.NewFileTypeUnknownError(findErr.Error(), "falha ao buscar tipo de ficheiro")
	}

	return fileType, nil
}

func (s *FileTypeService) UpdateFileType(fileTypeId string, name *string, description *string) *common.ErrorResponse {

	fileType, err := s.FindFileTypeById(fileTypeId)

	if err != nil {
		return err
	}

	fileType.Update(name, description)

	updatedErr := s.FileTypeRepo.Update(fileType)

	if updatedErr != nil {
		return errors.NewFileTypeUnknownError(updatedErr.Error(), "falha ao atualizar tipo de ficheiro")
	}

	return nil
}

func (s *FileTypeService) DeleteFileType(fileTypeId string) *common.ErrorResponse {

	fileType, err := s.FindFileTypeById(fileTypeId)

	if err != nil {
		return err
	}

	deletedErr := s.FileTypeRepo.Delete(fileType)

	if deletedErr != nil {
		return errors.NewFileTypeUnknownError(deletedErr.Error(), "falha ao deletar tipo de ficheiro")
	}

	return nil
}
