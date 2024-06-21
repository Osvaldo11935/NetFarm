package requests

import "time"

type UpdatePersonRequest struct {
	FullName  *string    `json:"fullName"`
	BirthDate *time.Time `json:"birthDate"`
	Gender    *string    `json:"gender"`
}
