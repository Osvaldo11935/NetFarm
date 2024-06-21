package services

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/errors"
	"NetFarm/shared/models/common"
	"time"
)

type PersonService struct {
	PersonRepo irepositories.IPersonRepository
}

func NewPersonService(PersonRepo irepositories.IPersonRepository) iservices.IPersonService {
	return &PersonService{
		PersonRepo: PersonRepo,
	}
}

func (s *PersonService) CreatePerson(Person *entities.Person) *common.ErrorResponse {

	p, _ := s.FindPersonByUserId(Person.UserId)

	if p != nil {
		return errors.NewPersonAlreadyExists("falha ao cadastrar dados pessoais do usuario")
	}

	createdErr := s.PersonRepo.Create(Person)

	if createdErr != nil {
		return errors.NewPersonUnknownError(createdErr.Error(), "falha ao cadastrar dados pessoais do usuario")
	}

	return nil
}

func (s *PersonService) FindPersonByUserId(userId string) (*entities.Person, *common.ErrorResponse) {
	var person entities.Person

	err := s.PersonRepo.Query().First(&person, "UserId", userId).Error

	if err != nil {
		return nil, errors.NewPersonUnknownError(err.Error(), "falha ao buscar dados pessoais do usuario")
	}

	return &person, nil
}

func (s *PersonService) UpdatePerson(userId string, fullName *string, birthDate *time.Time, gender *string) *common.ErrorResponse {

	Person, err := s.FindPersonByUserId(userId)

	if err != nil {
		return err
	}

	Person.Update(fullName, birthDate, gender)

	updatedErr := s.PersonRepo.Update(Person)

	if updatedErr != nil {
		return errors.NewPersonUnknownError(updatedErr.Error(), "falha ao atualizar dados pessoais do usuario")
	}

	return nil
}

func (s *PersonService) DeletePerson(userId string) *common.ErrorResponse {

	Person, err := s.FindPersonByUserId(userId)

	if err != nil {
		return err
	}

	deletedErr := s.PersonRepo.Delete(Person)

	if deletedErr != nil {
		return errors.NewPersonUnknownError(deletedErr.Error(), "falha ao deletar dados pessoais do usuario")
	}

	return nil
}
