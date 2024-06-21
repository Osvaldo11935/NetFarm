package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type CategoryService struct {
	CategoryRepo irepositories.ICategoryRepository
}

func NewCategoryService(CategoryRepo irepositories.ICategoryRepository) iservices.ICategoryService {
	return &CategoryService{
		CategoryRepo:CategoryRepo,
	}
}

func (s *CategoryService) CreateCategory(Category *entities.Category) *common.ErrorResponse {

	createdErr := s.CategoryRepo.Create(Category)

	if createdErr != nil {
		return errors.NewCategoryUnknownError(createdErr.Error(), "falha ao cadastrar estado")
	}

	return nil
}

func (s *CategoryService) FindAllCategory(pagination *common.Pagination) ([]entities.Category, *common.ErrorResponse) {
	var Category []entities.Category

	query := s.CategoryRepo.Query()

	query = query.Scopes(extensions.Paginate(&Category, pagination, query))

	findErr := query.Find(&Category).Error

	if findErr != nil {
		return Category, errors.NewCategoryUnknownError(findErr.Error(), "falha ao buscar estado")
	}

	pagination.Data =Category

	return Category, nil
}

func (s *CategoryService) FindBycategoryId(CategoryId string) (entities.Category, *common.ErrorResponse) {

	var Category entities.Category

	findErr := s.CategoryRepo.Query().First(&Category, "Id",CategoryId).Error

	if findErr != nil {
		return Category, errors.NewCategoryUnknownError(findErr.Error(), "falha ao buscar estado")
	}

	return Category, nil
}

func (s *CategoryService) UpdateCategory(categoryId string, name *string, description *string) *common.ErrorResponse {

	Category, err := s.FindBycategoryId(categoryId)

	if err != nil {
		return err
	}

	Category.Update(name, description)

	updatedErr := s.CategoryRepo.Update(Category)

	if updatedErr != nil {
		return errors.NewCategoryUnknownError(updatedErr.Error(), "falha ao atualizar estado")
	}

	return nil
}

func (s *CategoryService) DeleteCategory(CategoryId string) *common.ErrorResponse {

	Category, err := s.FindBycategoryId(CategoryId)

	if err != nil {
		return err
	}

	deletedErr := s.CategoryRepo.Delete(Category)

	if deletedErr != nil {
		return errors.NewCategoryUnknownError(deletedErr.Error(), "falha ao deletar estado")
	}

	return nil
}
