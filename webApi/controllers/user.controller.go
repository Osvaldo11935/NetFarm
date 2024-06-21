package controllers

import (
	"NetFarm/application/interfaces/iservices"
	"NetFarm/domain/entities"
	"NetFarm/shared/constants"
	"NetFarm/shared/models/common"
	"NetFarm/webApi/models/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type UserController struct {
	UserService iservices.IUserService
}

func NewUserController(service iservices.IUserService) *UserController {
	return &UserController{
		UserService: service,
	}
}

func (controller *UserController) HandleSignIn(ctx *gin.Context) {

	signInRequest := requests.SignInRequest{}
	err := ctx.ShouldBindJSON(&signInRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.UserService.SignIn(signInRequest.Email, signInRequest.Password)

	//if createErr != nil {
	//	ctx.JSON(http.StatusBadRequest, createErr)
	//	return
	//}

	//webResponse := common.BaseCreateResponse{Id: createUserRequest.Id}

	ctx.JSON(http.StatusOK, createErr)
}

func (controller *UserController) HandleCreateAdmin(ctx *gin.Context) {

	createUserRequest := entities.User{}
	err := ctx.ShouldBindJSON(&createUserRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.UserService.CreateUser(&createUserRequest, constants.Admin)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createUserRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) HandleCreateClient(ctx *gin.Context) {

	createUserRequest := entities.User{}
	err := ctx.ShouldBindJSON(&createUserRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.UserService.CreateUser(&createUserRequest, constants.Client)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createUserRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) HandleCreateEmployee(ctx *gin.Context) {

	createUserRequest := entities.User{}
	err := ctx.ShouldBindJSON(&createUserRequest)

	if err != nil {
		fmt.Println(err)
	}

	createErr := controller.UserService.CreateUser(&createUserRequest, constants.Employee)

	if createErr != nil {
		ctx.JSON(http.StatusBadRequest, createErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: createUserRequest.Id}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) HandleFindAllUsers(ctx *gin.Context) {

	page, _ := strconv.Atoi(ctx.Query("pageNumber"))
	pageSize, _ := strconv.Atoi(ctx.Query("pageSize"))

	pagination := common.Pagination{PageSize: pageSize, PageNumber: page}

	_, findErr := controller.UserService.FindAllUser(&pagination)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, pagination)
}

func (controller *UserController) HandleFindUserById(ctx *gin.Context) {
	userId := ctx.Param("userId")

	User, findErr := controller.UserService.FindUserById(userId)

	if findErr != nil {
		ctx.JSON(http.StatusBadRequest, findErr)
		return
	}

	ctx.JSON(http.StatusOK, User)
}

func (controller *UserController) HandleUpdateUser(ctx *gin.Context) {

	userId := ctx.Param("userId")

	createUserRequest := requests.UpdateUserRequest{}

	err := ctx.ShouldBindJSON(&createUserRequest)

	if err != nil {
		fmt.Println(err)
	}

	updatedErr := controller.UserService.UpdateUser(
		userId,
		createUserRequest.Email,
		createUserRequest.UserName,
		createUserRequest.PhoneNumber)

	if updatedErr != nil {
		ctx.JSON(http.StatusBadRequest, updatedErr)
		return
	}

	webResponse := common.BaseCreateResponse{Id: userId}

	ctx.JSON(http.StatusOK, webResponse)
}

func (controller *UserController) HandleDeleteUser(ctx *gin.Context) {

	userId := ctx.Param("userId")

	deletedErr := controller.UserService.DeleteUser(userId)

	if deletedErr != nil {
		ctx.JSON(http.StatusBadRequest, deletedErr)
		return
	}

	ctx.JSON(http.StatusNoContent, nil)
}
