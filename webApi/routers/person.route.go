package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddPersonRoute(route *gin.RouterGroup, controller *controllers.PersonController) {
	personRouter := route.Group("/person")
	personRouter.GET("user/:userId", controller.HandleFindPersonsByUserId)
	personRouter.POST("user/:userId", controller.HandleCreate)
	personRouter.PUT("user/:userId", controller.HandleUpdatePerson)
	personRouter.DELETE("user/:userId", controller.HandleDeletePerson)
}
