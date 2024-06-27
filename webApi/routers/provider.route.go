package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddProviderRoute(route *gin.RouterGroup, controller *controllers.ProviderController) {
	providerRouter := route.Group("/providers")
	providerRouter.GET("", controller.HandleFindAllProvider)
	providerRouter.POST("", controller.HandleCreate)
	providerRouter.GET(":providerId", controller.HandleFindProviderById)
	providerRouter.PUT(":providerId", controller.HandleUpdateProvider)
	providerRouter.DELETE(":providerId", controller.HandleDeleteProvider)
}
