package controllers

import (
	"NetFarm/application/interfaces/iinfrastructure/iexternalServices"
	"NetFarm/shared/models/requests"
	"NetFarm/shared/models/responses"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PaymentController struct {
	PaymentService iexternalServices.IPaymentService
}

func NewPaymentController(service iexternalServices.IPaymentService) *PaymentController {
	return &PaymentController{
		PaymentService: service,
	}
}

func (controller *PaymentController) HandleGetPaymentReference(ctx *gin.Context) {

	createPaymentRequest := requests.GetPaymentReferenceRequest{}
	err := ctx.ShouldBindJSON(&createPaymentRequest)

	if err != nil {
		fmt.Println(err)
	}

	reference, expTime := controller.PaymentService.GetPaymentReference(createPaymentRequest.Amount, createPaymentRequest.PhoneNumber, 1)

	webResponse := responses.GetPaymentReferenceResponse{Reference: reference, ExpirationTime: expTime}

	ctx.JSON(http.StatusOK, webResponse)
}

//func (controller *PaymentController) HandleSimulatePayment(ctx *gin.Context) {
//
//	_, createErr := controller.PaymentService.SimulatePayment()
//
//	if createErr != nil {
//		fmt.Println(createErr)
//	}
//
//	ctx.JSON(http.StatusOK, pagination)
//}
