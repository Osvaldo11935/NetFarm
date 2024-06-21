package controllers

import (
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
	"NetFarm/webApi/models/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PersonController struct {
	PersonService iservices.IPersonService
}

func NewPersonController(service iservices.IPersonService) *PersonController {
	return &PersonController{
		PersonService: service,
	}
}

func (controller *PersonController) HandleCreate(ctx *gin.Context) {

	createPersonRequest := entities.Person{}
	err := ctx.ShouldBindJSON(&createPersonRequest)

	if err != nil {
		fmt.Println(err)
	}

	userId := ctx.Param("userId")

	createPersonRequest.UserId = userId

	createErr := controller.PersonService.CreatePerson(&createPersonRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createPersonRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PersonController) HandleFindPersonsByUserId(ctx *gin.Context) {

	userId := ctx.Param("userId")

	persons, findErr := controller.PersonService.FindPersonByUserId(userId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, persons)
}

func (controller *PersonController) HandleFindPersonById(ctx *gin.Context) {
	personId := ctx.Param("personId")

	person, findErr := controller.PersonService.FindPersonByUserId(personId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, person)
}

func (controller *PersonController) HandleUpdatePerson(ctx *gin.Context) {

	personId := ctx.Param("personId")

	createPersonRequest := requests.UpdatePersonRequest{}

	err := ctx.ShouldBindJSON(&createPersonRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.PersonService.UpdatePerson(
		personId,
		createPersonRequest.FullName,
		createPersonRequest.BirthDate,
		createPersonRequest.Gender)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: personId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *PersonController) HandleDeletePerson(ctx *gin.Context) {

	personId := ctx.Param("personId")

	deletedErr := controller.PersonService.DeletePerson(personId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
