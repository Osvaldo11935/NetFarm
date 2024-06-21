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

type ImposedController struct {
	ImposedService iservices.IImposedService
}

func NewImposedController(service iservices.IImposedService) *ImposedController {
	return &ImposedController{
		ImposedService: service,
	}
}

func (controller *ImposedController) HandleCreate(ctx *gin.Context) {

	createImposedRequest := entities.Imposed{}
	err := ctx.ShouldBindJSON(&createImposedRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.ImposedService.CreateImposed(&createImposedRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createImposedRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ImposedController) HandleFindAllImposed(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.ImposedService.FindAllImposed(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *ImposedController) HandleFindImposedById(ctx *gin.Context) {
	imposedId := ctx.Param("imposedId")

	Imposed, findErr := controller.ImposedService.FindImposedById(imposedId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, Imposed)
}

func (controller *ImposedController) HandleUpdateImposed(ctx *gin.Context) {

	ImposedId := ctx.Param("ImposedId")

	createImposedRequest := requests.UpdateImposedRequest{}

	err := ctx.ShouldBindJSON(&createImposedRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.ImposedService.UpdateImposed(ImposedId, createImposedRequest.Name, createImposedRequest.Rate)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: ImposedId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ImposedController) HandleDeleteImposed(ctx *gin.Context) {

	imposedId := ctx.Param("imposedId")

	deletedErr := controller.ImposedService.DeleteImposed(imposedId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
