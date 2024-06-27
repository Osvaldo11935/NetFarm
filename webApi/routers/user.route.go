package routers

import (
	"NetFarm/webApi/controllers"
	"NetFarm/webApi/middleware"
	"github.com/gin-gonic/gin"
)

func AddUserRoute(route *gin.RouterGroup, controller *controllers.UserController) {
	userRouter := route.Group("/users")
	userRouter.GET("", controller.HandleFindAllUsers)
	userRouter.POST("sign-in", controller.HandleSignIn)
	userRouter.POST("admin", controller.HandleCreateAdmin)
	userRouter.POST("client", controller.HandleCreateClient)
	userRouter.POST("employee", controller.HandleCreateEmployee)
	userRouter.GET(":userId", controller.HandleFindUserById)
	userRouter.GET("user-info", middleware.AuthMiddleware(), controller.HandleFindUserInfo)
	userRouter.PUT(":userId", controller.HandleUpdateUser)
	userRouter.DELETE(":userId", controller.HandleDeleteUser)

}
