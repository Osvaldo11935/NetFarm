package requests

type UpdateUserRequest struct {
	Email       *string `json:"email"`
	UserName    *string `json:"userName"`
	PhoneNumber *string `json:"phoneNumber"`
}
