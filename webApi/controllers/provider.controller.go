package controllers

import (
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/models/common"
	"NetFarm/webApi/models/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type ProviderController struct {
	ProviderService iservices.IProviderService
}

func NewProviderController(ProviderService iservices.IProviderService) *ProviderController {
	return &ProviderController{
		ProviderService: ProviderService,
	}
}

func (controller *ProviderController) HandleCreate(ctx *gin.Context) {
	createProviderRequest := entities.Provider{}
	err := ctx.ShouldBindJSON(&createProviderRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.ProviderService.Create(&createProviderRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createProviderRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProviderController) HandleFindAllProvider(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.ProviderService.FindAllProvider(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *ProviderController) HandleFindProviderById(ctx *gin.Context) {
	providerId := ctx.Param("providerId")

	Provider, findErr := controller.ProviderService.FindProviderById(providerId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, Provider)
}

func (controller *ProviderController) HandleUpdateProvider(ctx *gin.Context) {

	providerId := ctx.Param("providerId")

	createProviderRequest := requests.UpdateProviderRequest{}

	err := ctx.ShouldBindJSON(&createProviderRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.ProviderService.UpdateProvider(
		providerId,
		createProviderRequest.Name)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: providerId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ProviderController) HandleDeleteProvider(ctx *gin.Context) {

	providerId := ctx.Param("providerId")

	deletedErr := controller.ProviderService.DeleteProvider(providerId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
