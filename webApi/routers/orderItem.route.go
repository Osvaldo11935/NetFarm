package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddOrderItemRoute(route *gin.RouterGroup, controller *controllers.OrderItemController) {
	ordersItemRouter := route.Group("/order-item")
	ordersItemRouter.GET("order/:orderId", controller.HandleFindOrderItemByOrderId)
	ordersItemRouter.POST("order/:orderId", controller.HandleCreate)
	ordersItemRouter.GET(":orderItemId", controller.HandleFindOrderItemById)
	ordersItemRouter.PUT(":orderItemId", controller.HandleUpdateOrderItem)
	ordersItemRouter.DELETE(":orderItemId", controller.HandleDeleteOrderItem)
}
