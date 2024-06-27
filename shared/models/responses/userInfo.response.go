package responses

import "NetFarm/domain/entities"

type UserInfoResponse struct {
	Id          string `json:"id"`
	Email       string `json:"email"`
	UserName    string `json:"userName"`
	PhoneNumber string `json:"phoneNumber"`
	Role        string `json:"role"`
}

func NewUserInfoResponse(user entities.User, role entities.Role) UserInfoResponse {
	return UserInfoResponse{
		Id:          user.Id,
		Email:       user.Email,
		UserName:    user.UserName,
		PhoneNumber: user.PhoneNumber,
		Role:        role.Name,
	}
}
