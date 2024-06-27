package main

import (
	"NetFarm/application/services"
	"NetFarm/infrastructure/externalService"
	"NetFarm/infrastructure/storages"
	"NetFarm/persistence/repositories"
	"NetFarm/webApi/controllers"
	"NetFarm/webApi/routers/common"
	"fmt"
	"net/http"
)

func main() {

	// Repository
	ordersRepo := repositories.NewOrderRepository()
	orderStatusRepo := repositories.NewOrderStatusRepository()
	orderFileRepo := repositories.NewOrderFileRepository()
	orderItemRepo := repositories.NewOrderItemRepository()
	fileTypeRepo := repositories.NewFileTypeRepository()
	roleRepo := repositories.NewRoleRepository()
	personRepo := repositories.NewPersonRepository()
	userRepo := repositories.NewUserRepository()
	addressRepo := repositories.NewAddressRepository()
	medicineRepo := repositories.NewMedicineRepository()
	providerRepo := repositories.NewProviderRepository()
	imposedRepo := repositories.NewImposedRepository()
	expenseRepo := repositories.NewExpenseRepository()
	orderDetailRepo := repositories.NewOrderDetailRepository()
	categoryRepo := repositories.NewCategoryRepository()

	// Service
	orderService := services.NewOrderService(ordersRepo)
	orderStatusService := services.NewOrderStatusService(orderStatusRepo)
	orderFileService := services.NewOrderFileService(orderFileRepo)
	orderItemService := services.NewOrderItemService(orderItemRepo)
	fileTypeService := services.NewFileTypeService(fileTypeRepo)
	roleService := services.NewRoleService(roleRepo)
	personService := services.NewPersonService(personRepo)
	addressService := services.NewAddressService(addressRepo)
	medicineService := services.NewMedicineService(medicineRepo)
	providerService := services.NewProviderService(providerRepo)
	paymentService := externalService.NewPaymentService()
	expenseService := services.NewExpenseService(expenseRepo)
	imposedService := services.NewImposedService(imposedRepo)
	orderDetailService := services.NewOrderDetailService(orderDetailRepo)
	jwtTokenService := externalService.NewJwtTokenService()
	userService := services.NewUserService(userRepo, jwtTokenService, roleService)
	categoryService := services.NewCategoryService(categoryRepo)

	// Storage
	fileManagerStorage := storages.NewGoogleStorage()

	// Controller
	orderController := controllers.NewOrderController(orderService)
	orderStatusController := controllers.NewOrderStatusController(orderStatusService)
	orderFileController := controllers.NewOrderFileController(orderFileService, fileManagerStorage)
	orderItemController := controllers.NewOrderItemController(orderItemService)
	fileTypeController := controllers.NewFileTypeController(fileTypeService)
	addressController := controllers.NewAddressController(addressService)
	roleController := controllers.NewRoleController(roleService)
	personController := controllers.NewPersonController(personService)
	userController := controllers.NewUserController(userService)
	providerController := controllers.NewProviderController(providerService)
	medicineController := controllers.NewMedicineController(medicineService)
	paymentController := controllers.NewPaymentController(paymentService)
	imposedController := controllers.NewImposedController(imposedService)
	expenseController := controllers.NewExpenseController(expenseService)
	orderDetailController := controllers.NewOrderDetailController(orderDetailService)
	categoryController := controllers.NewCategoryController(categoryService)

	// Router
	routes := common.NewRouter(
		orderController,
		orderFileController,
		orderStatusController,
		orderItemController,
		addressController,
		fileTypeController,
		personController,
		roleController,
		userController,
		medicineController,
		providerController,
		paymentController,
		expenseController,
		imposedController,
		orderDetailController,
		categoryController)

	server := &http.Server{
		Addr:    ":8000",
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
