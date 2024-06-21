package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	ordersController *controllers.OrderController, ordersFileController *controllers.OrderFileController,
	ordersStatusController *controllers.OrderStatusController, ordersItemController *controllers.OrderItemController,
	addressController *controllers.AddressController, fileTypeController *controllers.FileTypeController, personController *controllers.PersonController,
	roleController *controllers.RoleController, userController *controllers.UserController, medicineController *controllers.MedicineController,
	providerController *controllers.ProviderController, paymentController *controllers.PaymentController,
	expenseController *controllers.ExpenseController, imposedController *controllers.ImposedController, orderDetailController *controllers.OrderDetailController) *gin.Engine {

	router := gin.Default()

	//router.Use(middleware.SetHeaderMiddleware())
	baseRouter := router.Group("/api")

	//Route orders
	ordersRouter := baseRouter.Group("/orders")
	ordersRouter.GET("", ordersController.HandleFindAllOrders)
	ordersRouter.POST("user/:userId", ordersController.HandleCreate)
	ordersRouter.GET("user/:userId", ordersController.HandleFindOrdersByUserId)
	ordersRouter.GET("order/:orderId", ordersController.HandleFindOrderById)
	ordersRouter.PUT("order/:orderId", ordersController.HandleUpdateStatusOrder)
	ordersRouter.DELETE("order/:orderId", ordersController.HandleDeleteOrder)

	//Route orders status
	ordersStatusRouter := baseRouter.Group("/order-status")
	ordersStatusRouter.GET("", ordersStatusController.HandleFindAllOrderStatus)
	ordersStatusRouter.POST("", ordersStatusController.HandleCreate)
	ordersStatusRouter.GET(":orderStatusId", ordersStatusController.HandleFindOrderStatusById)
	ordersStatusRouter.PUT(":orderStatusId", ordersStatusController.HandleUpdateOrderStatus)
	ordersStatusRouter.DELETE(":orderStatusId", ordersStatusController.HandleDeleteOrderStatus)

	//Route orders item
	ordersItemRouter := baseRouter.Group("/order-item")
	ordersItemRouter.GET("order/:orderId", ordersItemController.HandleFindOrderItemByOrderId)
	ordersItemRouter.POST("order/:orderId", ordersItemController.HandleCreate)
	ordersItemRouter.GET(":orderItemId", ordersItemController.HandleFindOrderItemById)
	ordersItemRouter.PUT(":orderItemId", ordersItemController.HandleUpdateOrderItem)
	ordersItemRouter.DELETE(":orderItemId", ordersItemController.HandleDeleteOrderItem)

	//Route orders file
	ordersFileRouter := baseRouter.Group("/order-file")
	ordersFileRouter.GET("order/:orderId", ordersFileController.HandleFindAllOrderFiles)
	ordersFileRouter.POST("order/:orderId", ordersFileController.HandleCreate)
	ordersFileRouter.GET(":orderFileId", ordersFileController.HandleFindOrderFileById)
	ordersFileRouter.DELETE(":orderFileId", ordersFileController.HandleDeleteOrderFile)

	//Route address
	addressRouter := baseRouter.Group("/addresses")
	addressRouter.GET("user/:userId", addressController.HandleFindAddress)
	addressRouter.POST("user/:userId", addressController.HandleCreate)
	addressRouter.PUT("user/:userId", addressController.HandleUpdateAddress)
	addressRouter.DELETE("user/:userId", addressController.HandleDeleteAddress)

	//Route file type
	fileTypeRouter := baseRouter.Group("/file-type")
	fileTypeRouter.GET("", fileTypeController.HandleFindAllFileTypes)
	fileTypeRouter.GET(":fileTypeId", fileTypeController.HandleFindFileTypeById)
	fileTypeRouter.POST("", fileTypeController.HandleCreate)
	fileTypeRouter.PUT(":fileTypeId", fileTypeController.HandleUpdateFileType)
	fileTypeRouter.DELETE(":fileTypeId", fileTypeController.HandleDeleteFileType)

	//Route person
	personRouter := baseRouter.Group("/person")
	personRouter.GET("user/:userId", personController.HandleFindPersonsByUserId)
	personRouter.POST("user/:userId", personController.HandleCreate)
	personRouter.PUT("user/:userId", personController.HandleUpdatePerson)
	personRouter.DELETE("user/:userId", personController.HandleDeletePerson)

	//Route role
	roleRouter := baseRouter.Group("/roles")
	roleRouter.GET("", roleController.HandleFindAllRoles)
	roleRouter.POST("", roleController.HandleCreate)
	roleRouter.GET(":roleId", roleController.HandleFindRoleById)
	roleRouter.PUT(":roleId", roleController.HandleUpdateRole)
	roleRouter.DELETE(":roleId", roleController.HandleDeleteRole)

	//Route user
	userRouter := baseRouter.Group("/users")
	userRouter.GET("", userController.HandleFindAllUsers)
	userRouter.POST("sign-in", userController.HandleSignIn)
	userRouter.POST("admin", userController.HandleCreateAdmin)
	userRouter.POST("client", userController.HandleCreateClient)
	userRouter.POST("employee", userController.HandleCreateEmployee)
	userRouter.GET(":userId", userController.HandleFindUserById)
	userRouter.PUT(":userId", userController.HandleUpdateUser)
	userRouter.DELETE(":userId", userController.HandleDeleteUser)

	//Route medicine
	medicineRouter := baseRouter.Group("/medicines")
	medicineRouter.GET("", medicineController.HandleFindAllMedicine)
	medicineRouter.GET("provider/:providerId", medicineController.HandleFindMedicineByProviderId)
	medicineRouter.POST("provider/:providerId", medicineController.HandleCreate)
	medicineRouter.GET(":medicineId", medicineController.HandleFindMedicineById)
	medicineRouter.PUT(":medicineId", medicineController.HandleUpdateMedicine)
	medicineRouter.DELETE(":medicineId", medicineController.HandleDeleteMedicine)

	//Route provider
	providerRouter := baseRouter.Group("/providers")
	providerRouter.GET("", providerController.HandleFindAllProvider)
	providerRouter.POST("", providerController.HandleCreate)
	providerRouter.GET(":providerId", providerController.HandleFindProviderById)
	providerRouter.PUT(":providerId", providerController.HandleUpdateProvider)
	providerRouter.DELETE(":providerId", providerController.HandleDeleteProvider)

	//Route Imposed
	imposedRouter := baseRouter.Group("/imposed")
	imposedRouter.GET("", imposedController.HandleFindAllImposed)
	imposedRouter.POST("", imposedController.HandleCreate)
	imposedRouter.GET(":imposedId", imposedController.HandleFindImposedById)
	imposedRouter.PUT(":imposedId", imposedController.HandleUpdateImposed)
	imposedRouter.DELETE(":imposedId", imposedController.HandleDeleteImposed)

	//Route expense
	expenseRouter := baseRouter.Group("/expense")
	expenseRouter.GET("", expenseController.HandleFindAllExpense)
	expenseRouter.POST("", expenseController.HandleCreate)
	expenseRouter.GET(":expenseId", expenseController.HandleFindExpenseById)
	expenseRouter.PUT(":expenseId", expenseController.HandleUpdateExpense)
	expenseRouter.DELETE(":expenseId", expenseController.HandleDeleteExpense)

	//Route order detail
	orderDetailRouter := baseRouter.Group("/order-detail")
	orderDetailRouter.GET("", orderDetailController.HandleFindAllOrderDetails)
	orderDetailRouter.POST("", orderDetailController.HandleCreate)
	orderDetailRouter.POST("calculate", orderDetailController.HandleOrderCalculate)

	//Route payment
	paymentRouter := baseRouter.Group("/payment")
	paymentRouter.POST("/reference", paymentController.HandleGetPaymentReference)
	return router
}
