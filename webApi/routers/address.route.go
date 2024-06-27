package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddAddressRoute(router *gin.RouterGroup, controller *controllers.AddressController) {
	addressRouter := router.Group("/addresses")
	addressRouter.GET("user/:userId", controller.HandleFindAddress)
	addressRouter.POST("user/:userId", controller.HandleCreate)
	addressRouter.PUT("user/:userId", controller.HandleUpdateAddress)
	addressRouter.DELETE("user/:userId", controller.HandleDeleteAddress)
}
