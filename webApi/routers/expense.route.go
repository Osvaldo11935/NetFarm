package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddExpenseRoute(route *gin.RouterGroup, controller *controllers.ExpenseController) {
	expenseRouter := route.Group("/expense")
	expenseRouter.GET("", controller.HandleFindAllExpense)
	expenseRouter.POST("", controller.HandleCreate)
	expenseRouter.GET(":expenseId", controller.HandleFindExpenseById)
	expenseRouter.PUT(":expenseId", controller.HandleUpdateExpense)
	expenseRouter.DELETE(":expenseId", controller.HandleDeleteExpense)
}
