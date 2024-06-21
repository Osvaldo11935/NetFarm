package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type RoleService struct {
	RoleRepo irepositories.IRoleRepository
}

func NewRoleService(RoleRepo irepositories.IRoleRepository) iservices.IRoleService {
	return &RoleService{
		RoleRepo: RoleRepo,
	}
}

func (s *RoleService) CreateRole(Role *entities.Role) *common.ErrorResponse {

	createdErr := s.RoleRepo.Create(Role)

	if createdErr != nil {
		return errors.NewRoleUnknownError(createdErr.Error(), "falha ao cadastrar perfil")
	}

	return nil
}

func (s *RoleService) FindAllRole(pagination *common.Pagination) ([]entities.Role, *common.ErrorResponse) {
	var roles []entities.Role

	query := s.RoleRepo.Query()

	query = query.Scopes(extensions.Paginate(&roles, pagination, query))

	err := query.Find(&roles).Error

	if err != nil {
		return roles, errors.NewRoleUnknownError(err.Error(), "falha ao buscar perfil")
	}

	pagination.Data = roles

	return roles, nil
}

func (s *RoleService) FindRoleById(roleId string) (entities.Role, *common.ErrorResponse) {

	var role entities.Role

	err := s.RoleRepo.Query().First(&role, "Id", roleId).Error

	if err != nil {
		return role, errors.NewRoleUnknownError(err.Error(), "falha ao buscar perfil")
	}

	return role, nil
}

func (s *RoleService) UpdateRole(roleId string, name *string) *common.ErrorResponse {

	role, findErr := s.FindRoleById(roleId)

	if findErr != nil {
		return findErr
	}

	role.Update(name)

	updatedErr := s.RoleRepo.Update(role)

	if updatedErr != nil {
		return errors.NewRoleUnknownError(updatedErr.Error(), "falha ao atualizar perfil")
	}

	return nil
}

func (s *RoleService) DeleteRole(roleId string) *common.ErrorResponse {

	role, err := s.FindRoleById(roleId)

	if err != nil {
		return err
	}

	deletedErr := s.RoleRepo.Delete(role)

	if deletedErr != nil {
		return errors.NewRoleUnknownError(deletedErr.Error(), "falha ao deletar perfil")
	}

	return nil
}
