package common

import (
	"NetFarm/webApi/controllers"
	"NetFarm/webApi/routers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"time"
)

func NewRouter(
	ordersController *controllers.OrderController, ordersFileController *controllers.OrderFileController,
	ordersStatusController *controllers.OrderStatusController, ordersItemController *controllers.OrderItemController,
	addressController *controllers.AddressController, fileTypeController *controllers.FileTypeController, personController *controllers.PersonController,
	roleController *controllers.RoleController, userController *controllers.UserController, medicineController *controllers.MedicineController,
	providerController *controllers.ProviderController, paymentController *controllers.PaymentController,
	expenseController *controllers.ExpenseController, imposedController *controllers.ImposedController, orderDetailController *controllers.OrderDetailController,
	categoryController *controllers.CategoryController) *gin.Engine {

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	//Add Route
	baseRouter := router.Group("/api")
	routers.AddOrderRoute(baseRouter, ordersController)
	routers.AddOrderStatusRoute(baseRouter, ordersStatusController)
	routers.AddOrderItemRoute(baseRouter, ordersItemController)
	routers.AddOrderFileRoute(baseRouter, ordersFileController)
	routers.AddAddressRoute(baseRouter, addressController)
	routers.AddFileTypeRoute(baseRouter, fileTypeController)
	routers.AddPersonRoute(baseRouter, personController)
	routers.AddRoleRoute(baseRouter, roleController)
	routers.AddCategoryRoute(baseRouter, categoryController)
	routers.AddUserRoute(baseRouter, userController)
	routers.AddMedicineRoute(baseRouter, medicineController)
	routers.AddProviderRoute(baseRouter, providerController)
	routers.AddImposedRoute(baseRouter, imposedController)
	routers.AddExpenseRoute(baseRouter, expenseController)
	routers.AddOrderDetailRoute(baseRouter, orderDetailController)
	routers.AddPaymentRoute(baseRouter, paymentController)

	return router
}
