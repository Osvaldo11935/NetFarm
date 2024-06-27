package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddOrderFileRoute(route *gin.RouterGroup, controller *controllers.OrderFileController) {
	ordersFileRouter := route.Group("/order-file")
	ordersFileRouter.GET("order/:orderId", controller.HandleFindAllOrderFiles)
	ordersFileRouter.POST("order/:orderId", controller.HandleCreate)
	ordersFileRouter.GET(":orderFileId", controller.HandleFindOrderFileById)
	ordersFileRouter.DELETE(":orderFileId", controller.HandleDeleteOrderFile)
}
