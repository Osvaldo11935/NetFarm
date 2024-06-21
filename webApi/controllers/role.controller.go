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

type RoleController struct {
	RoleService iservices.IRoleService
}

func NewRoleController(service iservices.IRoleService) *RoleController {
	return &RoleController{
		RoleService: service,
	}
}

func (controller *RoleController) HandleCreate(ctx *gin.Context) {

	createRoleRequest := entities.Role{}
	err := ctx.ShouldBindJSON(&createRoleRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.RoleService.CreateRole(&createRoleRequest)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createRoleRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RoleController) HandleFindAllRoles(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.RoleService.FindAllRole(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *RoleController) HandleFindRoleById(ctx *gin.Context) {
	roleId := ctx.Param("roleId")

	Role, findErr := controller.RoleService.FindRoleById(roleId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, Role)
}

func (controller *RoleController) HandleUpdateRole(ctx *gin.Context) {

	roleId := ctx.Param("roleId")

	createRoleRequest := requests.UpdateRoleRequest{}

	err := ctx.ShouldBindJSON(&createRoleRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.RoleService.UpdateRole(
		roleId,
		createRoleRequest.Name)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: roleId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *RoleController) HandleDeleteRole(ctx *gin.Context) {

	roleId := ctx.Param("roleId")

	deletedErr := controller.RoleService.DeleteRole(roleId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
