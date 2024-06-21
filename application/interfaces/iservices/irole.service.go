package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IRoleService interface {
	CreateRole(Role *entities.Role) *common.ErrorResponse
	FindAllRole(pagination *common.Pagination) ([]entities.Role, *common.ErrorResponse)
	FindRoleById(roleId string) (entities.Role, *common.ErrorResponse)
	UpdateRole(roleId string, name *string) *common.ErrorResponse
	DeleteRole(roleId string) *common.ErrorResponse
}
