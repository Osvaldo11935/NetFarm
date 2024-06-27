package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
	"NetFarm/shared/models/responses"
)

type IUserService interface {
	SignIn(email string, password string) (*string, *common.ErrorResponse)
	CreateUser(User *entities.User, roleName string) *common.ErrorResponse
	FindAllUser(pagination *common.Pagination) ([]entities.User, *common.ErrorResponse)
	FindUserById(UserId string) (entities.User, *common.ErrorResponse)
	FindUserInfo(userId string) (*responses.UserInfoResponse, *common.ErrorResponse)
	UpdateUser(UserId string, email *string, userName *string, phoneNumber *string) *common.ErrorResponse
	DeleteUser(UserId string) *common.ErrorResponse
}
