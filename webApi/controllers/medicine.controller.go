package controllers

import (
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	common "NetFarm/shared/models/common"
	"NetFarm/webApi/models/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type MedicineController struct {
	MedicineService iservices.IMedicineService
}

func NewMedicineController(MedicineService iservices.IMedicineService) *MedicineController {
	return &MedicineController{
		MedicineService: MedicineService,
	}
}

func (controller *MedicineController) HandleCreate(ctx *gin.Context) {
	providerId := ctx.Param("providerId")

	createMedicineRequest := entities.Medicine{}

	err := ctx.ShouldBindJSON(&createMedicineRequest)

	if err != nil {
		fmt.Println(err)
	}

	createMedicineRequest.ProviderId = providerId

	createErr := controller.MedicineService.Create(&createMedicineRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createMedicineRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MedicineController) HandleCreateMedicineCategory(ctx *gin.Context) {
	medicineId := ctx.Param("medicineId")

	var medicineCategoryIdsRequest []string

	err := ctx.ShouldBindJSON(&medicineCategoryIdsRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.MedicineService.CreateMedicineCategory(medicineId, medicineCategoryIdsRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}

func (controller *MedicineController) HandleFindAllMedicine(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.MedicineService.FindAllMedicine(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *MedicineController) HandleFindMedicineByProviderId(ctx *gin.Context) {
	providerId := ctx.Param("providerId")

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.MedicineService.FindMedicineByProvider(providerId, &pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *MedicineController) HandleFindMedicineById(ctx *gin.Context) {
	medicineId := ctx.Param("medicineId")

	Medicine, findErr := controller.MedicineService.FindMedicineById(medicineId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, Medicine)
}

func (controller *MedicineController) HandleUpdateMedicine(ctx *gin.Context) {

	medicineId := ctx.Param("medicineId")

	createMedicineRequest := requests.UpdateMedicineRequest{}

	err := ctx.ShouldBindJSON(&createMedicineRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.MedicineService.UpdateMedicine(
		medicineId,
		createMedicineRequest.Name,
		createMedicineRequest.Description,
		createMedicineRequest.Quantity,
		createMedicineRequest.Price)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: medicineId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *MedicineController) HandleDeleteMedicine(ctx *gin.Context) {

	medicineId := ctx.Param("medicineId")

	deletedErr := controller.MedicineService.DeleteMedicine(medicineId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
