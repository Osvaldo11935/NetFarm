package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddOrderStatusRoute(route *gin.RouterGroup, controller *controllers.OrderStatusController) {
	ordersStatusRouter := route.Group("/order-status")
	ordersStatusRouter.GET("", controller.HandleFindAllOrderStatus)
	ordersStatusRouter.POST("", controller.HandleCreate)
	ordersStatusRouter.GET(":orderStatusId", controller.HandleFindOrderStatusById)
	ordersStatusRouter.PUT(":orderStatusId", controller.HandleUpdateOrderStatus)
	ordersStatusRouter.DELETE(":orderStatusId", controller.HandleDeleteOrderStatus)
}
