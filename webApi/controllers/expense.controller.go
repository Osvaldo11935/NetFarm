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

type ExpenseController struct {
	ExpenseService iservices.IExpenseService
}

func NewExpenseController(service iservices.IExpenseService) *ExpenseController {
	return &ExpenseController{
		ExpenseService: service,
	}
}

func (controller *ExpenseController) HandleCreate(ctx *gin.Context) {

	createExpenseRequest := entities.Expense{}
	err := ctx.ShouldBindJSON(&createExpenseRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.ExpenseService.CreateExpense(&createExpenseRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createExpenseRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ExpenseController) HandleFindAllExpense(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.ExpenseService.FindAllExpense(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *ExpenseController) HandleFindExpenseById(ctx *gin.Context) {
	expenseId := ctx.Param("expenseId")

	Expense, findErr := controller.ExpenseService.FindExpenseById(expenseId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, Expense)
}

func (controller *ExpenseController) HandleUpdateExpense(ctx *gin.Context) {

	expenseId := ctx.Param("expenseId")

	createExpenseRequest := requests.UpdateExpenseRequest{}

	err := ctx.ShouldBindJSON(&createExpenseRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.ExpenseService.UpdateExpense(expenseId, createExpenseRequest.Name, createExpenseRequest.Rate)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: expenseId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *ExpenseController) HandleDeleteExpense(ctx *gin.Context) {

	expenseId := ctx.Param("expenseId")

	deletedErr := controller.ExpenseService.DeleteExpense(expenseId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
