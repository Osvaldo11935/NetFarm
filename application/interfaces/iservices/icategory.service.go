package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type ICategoryService interface {
	CreateCategory(Category *entities.Category) *common.ErrorResponse
	FindAllCategory(pagination *common.Pagination) ([]entities.Category, *common.ErrorResponse)
	FindBycategoryId(categoryId string) (entities.Category, *common.ErrorResponse)

	UpdateCategory(categoryId string, name string, description *string) *common.ErrorResponse

	DeleteCategory(categoryId string) *common.ErrorResponse
}
