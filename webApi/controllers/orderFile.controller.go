package controllers

import (
	"NetFarm/application/interfaces/iinfrastructure/istorages"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	common2 "NetFarm/shared/models/common"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type OrderFileController struct {
	OrderFileService   iservices.IOrderFileService
	FileManagerStorage istorages.IFileManagerStorage
}

func NewOrderFileController(orderFileService iservices.IOrderFileService,
	fileManagerStorage istorages.IFileManagerStorage) *OrderFileController {
	return &OrderFileController{
		OrderFileService:   orderFileService,
		FileManagerStorage: fileManagerStorage,
	}
}

func (controller *OrderFileController) HandleCreate(ctx *gin.Context) {
	createOrderFileRequest := entities.OrderFile{}

	orderId := ctx.Param("orderId")
	fileTypId := ctx.Param("fileTypId")

	createOrderFileRequest.OrderId = orderId
	createOrderFileRequest.FileTypeId = fileTypId

	url, uploadErr := controller.FileManagerStorage.Upload(ctx)

	if uploadErr != nil {
		ctx.JSON(http.StatusBadRequest, uploadErr)
		return
	}

	createOrderFileRequest.Url = url

	createErr := controller.OrderFileService.CreateOrderFile(&createOrderFileRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := map[string]interface{}{"id": createOrderFileRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *OrderFileController) HandleFindAllOrderFiles(ctx *gin.Context) {

	orderId := ctx.Param("orderId")

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common2.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.OrderFileService.FindOrderFileByOrderId(orderId, &pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *OrderFileController) HandleFindOrderFileById(ctx *gin.Context) {
	orderFileId := ctx.Param("orderFileId")

	orderFile, findErr := controller.OrderFileService.FindOrderFileById(orderFileId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, orderFile)
}

func (controller *OrderFileController) HandleDeleteOrderFile(ctx *gin.Context) {

	orderFileId := ctx.Param("orderFileId")

	deletedErr := controller.OrderFileService.DeleteOrderFile(orderFileId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
