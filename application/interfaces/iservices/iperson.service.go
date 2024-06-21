package iservices

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
	"time"
)

type IPersonService interface {
	CreatePerson(Person *entities.Person) *common.ErrorResponse
	FindPersonByUserId(userId string) (*entities.Person, *common.ErrorResponse)
	UpdatePerson(personId string, fullName *string, birthDate *time.Time, gender *string) *common.ErrorResponse
	DeletePerson(personId string) *common.ErrorResponse
}
