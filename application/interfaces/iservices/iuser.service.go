package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
)

type IUserService interface {
	SignIn(email string, password string) string
	CreateUser(User *entities.User, roleName string) *common.ErrorResponse
	FindAllUser(pagination *common.Pagination) ([]entities.User, *common.ErrorResponse)
	FindUserById(UserId string) (entities.User, *common.ErrorResponse)
	UpdateUser(UserId string, email *string, userName *string, phoneNumber *string) *common.ErrorResponse
	DeleteUser(UserId string) *common.ErrorResponse
}
