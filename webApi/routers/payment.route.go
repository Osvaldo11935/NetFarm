package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddPaymentRoute(route *gin.RouterGroup, controller *controllers.PaymentController) {
	paymentRouter := route.Group("/payment")
	paymentRouter.POST("/reference", controller.HandleGetPaymentReference)
}
