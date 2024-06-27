package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddRoleRoute(route *gin.RouterGroup, controller *controllers.RoleController) {
	roleRouter := route.Group("/roles")
	roleRouter.GET("", controller.HandleFindAllRoles)
	roleRouter.POST("", controller.HandleCreate)
	roleRouter.GET(":roleId", controller.HandleFindRoleById)
	roleRouter.PUT(":roleId", controller.HandleUpdateRole)
	roleRouter.DELETE(":roleId", controller.HandleDeleteRole)
}
