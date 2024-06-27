package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddCategoryRoute(route *gin.RouterGroup, controller *controllers.CategoryController) {
	categoryRouter := route.Group("/category")
	categoryRouter.GET("", controller.HandleFindAllCategory)
	categoryRouter.POST("", controller.HandleCreate)
	categoryRouter.GET(":categoryId", controller.HandleFindCategoryById)
	categoryRouter.PUT(":categoryId", controller.HandleUpdateCategory)
	categoryRouter.DELETE(":categoryId", controller.HandleDeleteCategory)

}
