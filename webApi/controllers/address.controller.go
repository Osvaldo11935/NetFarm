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

type AddressController struct {
	AddressService iservices.IAddressService
}

func NewAddressController(service iservices.IAddressService) *AddressController {
	return &AddressController{
		AddressService: service,
	}
}

func (controller *AddressController) HandleCreate(ctx *gin.Context) {

	userId := ctx.Param("userId")

	createAddressRequest := entities.Address{}

	unknownErr := ctx.ShouldBindJSON(&createAddressRequest)

	if unknownErr != nil {
		ctx.JSON(http.StatusBadRequest, unknownErr)
		return
	}

	createAddressRequest.UserId = userId

	createErr := controller.AddressService.CreateAddress(&createAddressRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createAddressRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AddressController) HandleFindAddress(ctx *gin.Context) {

	userId := ctx.Param("userId")

	addresses, findErr := controller.AddressService.FindAddressByUserId(userId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, addresses)
}

func (controller *AddressController) HandleUpdateAddress(ctx *gin.Context) {

	userId := ctx.Param("userId")

	createAddressRequest := requests.UpdateAddressRequest{}

	err := ctx.ShouldBindJSON(&createAddressRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.AddressService.UpdateAddress(
		userId,
		createAddressRequest.Country,
		createAddressRequest.State,
		createAddressRequest.City,
		createAddressRequest.Neighborhood,
		createAddressRequest.Street,
		createAddressRequest.Number,
		createAddressRequest.Complement)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: userId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *AddressController) HandleDeleteAddress(ctx *gin.Context) {

	userId := ctx.Param("userId")

	deletedErr := controller.AddressService.DeleteAddress(userId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
