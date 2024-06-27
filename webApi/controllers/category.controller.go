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

type CategoryController struct {
	CategoryService iservices.ICategoryService
}

func NewCategoryController(service iservices.ICategoryService) *CategoryController {
	return &CategoryController{
		CategoryService: service,
	}
}

func (controller *CategoryController) HandleCreate(ctx *gin.Context) {

	createCategoryRequest := entities.Category{}
	err := ctx.ShouldBindJSON(&createCategoryRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.CategoryService.CreateCategory(&createCategoryRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createCategoryRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryController) HandleFindAllCategory(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.CategoryService.FindAllCategory(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *CategoryController) HandleFindCategoryById(ctx *gin.Context) {
	categoryId := ctx.Param("categoryId")

	Category, findErr := controller.CategoryService.FindByCategoryId(categoryId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, Category)
}

func (controller *CategoryController) HandleUpdateCategory(ctx *gin.Context) {

	categoryId := ctx.Param("categoryId")

	createCategoryRequest := requests.UpdateCategoryRequest{}

	err := ctx.ShouldBindJSON(&createCategoryRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.CategoryService.UpdateCategory(
		categoryId,
		createCategoryRequest.Name,
		createCategoryRequest.Description)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: categoryId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *CategoryController) HandleDeleteCategory(ctx *gin.Context) {

	categoryId := ctx.Param("categoryId")

	deletedErr := controller.CategoryService.DeleteCategory(categoryId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
