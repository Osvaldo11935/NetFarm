package setup

import (
	"NetFarm/application/interfaces/iinfrastructure/istorages"
	"NetFarm/webApi/controllers"
)

type ControllersSetup struct {
	OrderController       *controllers.OrderController
	OrderStatusController *controllers.OrderStatusController
	OrderFileController   *controllers.OrderFileController
	OrderItemController   *controllers.OrderItemController
	FileTypeController    *controllers.FileTypeController
	AddressController     *controllers.AddressController
	RoleController        *controllers.RoleController
	PersonController      *controllers.PersonController
	UserController        *controllers.UserController
	ProviderController    *controllers.ProviderController
	MedicineController    *controllers.MedicineController
	PaymentController     *controllers.PaymentController
	ImposedController     *controllers.ImposedController
	ExpenseController     *controllers.ExpenseController
	OrderDetailController *controllers.OrderDetailController
	CategoryController    *controllers.CategoryController
}

func NewControllers(services *ServicesSetup, storage *istorages.IFileManagerStorage) *ControllersSetup {
	return &ControllersSetup{
		OrderController:       controllers.NewOrderController(services.OrderService),
		OrderStatusController: controllers.NewOrderStatusController(services.OrderStatusService),
		OrderFileController:   controllers.NewOrderFileController(services.OrderFileService, *storage),
		OrderItemController:   controllers.NewOrderItemController(services.OrderItemService),
		FileTypeController:    controllers.NewFileTypeController(services.FileTypeService),
		AddressController:     controllers.NewAddressController(services.AddressService),
		RoleController:        controllers.NewRoleController(services.RoleService),
		PersonController:      controllers.NewPersonController(services.PersonService),
		UserController:        controllers.NewUserController(services.UserService),
		ProviderController:    controllers.NewProviderController(services.ProviderService),
		MedicineController:    controllers.NewMedicineController(services.MedicineService),
		PaymentController:     controllers.NewPaymentController(services.PaymentService),
		ImposedController:     controllers.NewImposedController(services.ImposedService),
		ExpenseController:     controllers.NewExpenseController(services.ExpenseService),
		OrderDetailController: controllers.NewOrderDetailController(services.OrderDetailService),
		CategoryController:    controllers.NewCategoryController(services.CategoryService),
	}
}
