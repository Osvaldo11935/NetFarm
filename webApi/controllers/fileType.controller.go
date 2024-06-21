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

type FileTypeController struct {
	FileTypeService iservices.IFileTypeService
}

func NewFileTypeController(service iservices.IFileTypeService) *FileTypeController {
	return &FileTypeController{
		FileTypeService: service,
	}
}

func (controller *FileTypeController) HandleCreate(ctx *gin.Context) {

	createFileTypeRequest := entities.FileType{}
	err := ctx.ShouldBindJSON(&createFileTypeRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.FileTypeService.CreateFileType(&createFileTypeRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createFileTypeRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FileTypeController) HandleFindAllFileTypes(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.FileTypeService.FindAllFileType(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *FileTypeController) HandleFindFileTypeById(ctx *gin.Context) {
	fileTypeId := ctx.Param("fileTypeId")

	fileType, findErr := controller.FileTypeService.FindFileTypeById(fileTypeId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, fileType)
}

func (controller *FileTypeController) HandleUpdateFileType(ctx *gin.Context) {

	fileTypeId := ctx.Param("fileTypeId")

	createFileTypeRequest := requests.UpdateFileTypeRequest{}

	err := ctx.ShouldBindJSON(&createFileTypeRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.FileTypeService.UpdateFileType(fileTypeId, createFileTypeRequest.Type, createFileTypeRequest.Description)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: fileTypeId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *FileTypeController) HandleDeleteFileType(ctx *gin.Context) {

	fileTypeId := ctx.Param("fileTypeId")

	deletedErr := controller.FileTypeService.DeleteFileType(fileTypeId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
