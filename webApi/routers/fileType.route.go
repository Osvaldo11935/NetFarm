package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddFileTypeRoute(route *gin.RouterGroup, controller *controllers.FileTypeController) {
	fileTypeRouter := route.Group("/file-type")
	fileTypeRouter.GET("", controller.HandleFindAllFileTypes)
	fileTypeRouter.GET(":fileTypeId", controller.HandleFindFileTypeById)
	fileTypeRouter.POST("", controller.HandleCreate)
	fileTypeRouter.PUT(":fileTypeId", controller.HandleUpdateFileType)
	fileTypeRouter.DELETE(":fileTypeId", controller.HandleDeleteFileType)

}
