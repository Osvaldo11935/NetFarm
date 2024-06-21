package services

import (
	"NetFarm/application/extensions"
	"NetFarm/application/interfaces/iinfrastructure/iexternalServices"
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
)

type UserService struct {
	UserRepo     irepositories.IUserRepository
	RoleService  iservices.IRoleService
	TokenService iexternalServices.IJwtTokenService
}

func NewUserService(userRepo irepositories.IUserRepository, tokenService iexternalServices.IJwtTokenService,
	roleService iservices.IRoleService) iservices.IUserService {
	return &UserService{
		UserRepo:     userRepo,
		RoleService:  roleService,
		TokenService: tokenService,
	}
}

func (s *UserService) CreateUser(User *entities.User, roleName string) *common.ErrorResponse {

	role, findRoleErr := s.FindRoleByName(roleName, "falha ao cadastrar usuario")

	if findRoleErr != nil {
		return findRoleErr
	}

	User.RoleId = role.Id

	createdErr := s.UserRepo.Create(User)

	if createdErr != nil {
		return errors.NewUserUnknownError(createdErr.Error(), "falha ao cadastrar usuario")
	}

	return nil
}

func (s *UserService) SignIn(email string, password string) string {

	var claims []string

	var user entities.User

	query := s.UserRepo.Query().
		Where("Email", email)
	//Where("Password", password)

	findErr := query.First(&user).Error

	if findErr != nil {
		//"", errors.NewUserUnknownError(findErr.Error(), "falha ao buscar usuario")
	}

	role, findRoleErr := s.RoleService.FindRoleById(user.RoleId)

	if findRoleErr != nil {

	}

	claims = append(claims, role.Name)

	token, tokenErr := s.TokenService.GenerateToken(user.Email, user.Id, claims)

	if tokenErr != nil {

	}

	return token

}

func (s *UserService) FindAllUser(pagination *common.Pagination) ([]entities.User, *common.ErrorResponse) {
	var users []entities.User

	query := s.UserRepo.Query()

	query = query.Scopes(extensions.Paginate(&users, pagination, query))

	err := query.Find(&users).Error

	if err != nil {
		return users, errors.NewUserUnknownError(err.Error(), "falha ao buscar usuario")
	}

	pagination.Data = users

	return users, nil
}

func (s *UserService) FindUserById(UserId string) (entities.User, *common.ErrorResponse) {

	var User entities.User

	err := s.UserRepo.Query().First(&User, "Id", UserId).Error

	if err != nil {
		return User, errors.NewUserUnknownError(err.Error(), "falha ao buscar usuario")
	}

	return User, nil
}

func (s *UserService) UpdateUser(UserId string, email *string, userName *string, phoneNumber *string) *common.ErrorResponse {

	User, findErr := s.FindUserById(UserId)

	if findErr != nil {
		return findErr
	}

	User.Update(email, userName, phoneNumber)

	updateErr := s.UserRepo.Update(User)

	if updateErr != nil {
		return errors.NewUserUnknownError(updateErr.Error(), "falha ao atualizar endereço")
	}

	return nil
}

func (s *UserService) DeleteUser(UserId string) *common.ErrorResponse {

	User, err := s.FindUserById(UserId)

	if err != nil {
		return err
	}
	deletedErr := s.UserRepo.Delete(User)

	if deletedErr != nil {
		return errors.NewUserUnknownError(deletedErr.Error(), "falha ao deletar endereço")
	}
	return nil
}

func (s *UserService) FindRoleByName(name string, message string) (*entities.Role, *common.ErrorResponse) {

	var role entities.Role

	findErr := s.UserRepo.Query().First(&role, "Name", name).Error

	if findErr != nil {
		return nil, errors.NewRoleNotFound(message)
	}

	return &role, nil
}
