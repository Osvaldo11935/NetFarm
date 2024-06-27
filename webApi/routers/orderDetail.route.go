package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddOrderDetailRoute(route *gin.RouterGroup, controller *controllers.OrderDetailController) {
	orderDetailRouter := route.Group("/order-detail")
	orderDetailRouter.GET("", controller.HandleFindAllOrderDetails)
	orderDetailRouter.POST("", controller.HandleCreate)
	orderDetailRouter.POST("calculate", controller.HandleOrderCalculate)
}
