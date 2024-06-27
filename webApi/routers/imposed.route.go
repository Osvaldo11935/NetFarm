package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddImposedRoute(route *gin.RouterGroup, controller *controllers.ImposedController) {
	imposedRouter := route.Group("/imposed")
	imposedRouter.GET("", controller.HandleFindAllImposed)
	imposedRouter.POST("", controller.HandleCreate)
	imposedRouter.GET(":imposedId", controller.HandleFindImposedById)
	imposedRouter.PUT(":imposedId", controller.HandleUpdateImposed)
	imposedRouter.DELETE(":imposedId", controller.HandleDeleteImposed)

}
