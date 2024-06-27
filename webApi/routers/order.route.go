package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddOrderRoute(route *gin.RouterGroup, controller *controllers.OrderController) {
	ordersRouter := route.Group("/orders")
	ordersRouter.GET("", controller.HandleFindAllOrders)
	ordersRouter.POST("user/:userId", controller.HandleCreate)
	ordersRouter.GET("user/:userId", controller.HandleFindOrdersByUserId)
	ordersRouter.GET(":orderId", controller.HandleFindOrderById)
	ordersRouter.PUT(":orderId", controller.HandleUpdateStatusOrder)
	ordersRouter.DELETE(":orderId", controller.HandleDeleteOrder)
}
