package setup

import (
	"NetFarm/application/interfaces/iinfrastructure/iexternalServices"
	"NetFarm/application/interfaces/iservices"
	"NetFarm/application/services"
	"NetFarm/infrastructure/externalService"
)

type ServicesSetup struct {
	OrderService       iservices.IOrderService
	OrderStatusService iservices.IOrderStatusService
	OrderFileService   iservices.IOrderFileService
	OrderItemService   iservices.IOrderItemService
	FileTypeService    iservices.IFileTypeService
	RoleService        iservices.IRoleService
	PersonService      iservices.IPersonService
	AddressService     iservices.IAddressService
	MedicineService    iservices.IMedicineService
	ProviderService    iservices.IProviderService
	PaymentService     iexternalServices.IPaymentService
	ExpenseService     iservices.IExpenseService
	ImposedService     iservices.IImposedService
	OrderDetailService iservices.IOrderDetailService
	JwtTokenService    iexternalServices.IJwtTokenService
	UserService        iservices.IUserService
	CategoryService    iservices.ICategoryService
}

func NewServices(repos *RepositoriesSetup) *ServicesSetup {
	jwtTokenService := externalService.NewJwtTokenService()
	roleService := services.NewRoleService(repos.RoleRepo)

	return &ServicesSetup{
		OrderService:       services.NewOrderService(repos.OrdersRepo),
		OrderStatusService: services.NewOrderStatusService(repos.OrderStatusRepo),
		OrderFileService:   services.NewOrderFileService(repos.OrderFileRepo),
		OrderItemService:   services.NewOrderItemService(repos.OrderItemRepo),
		FileTypeService:    services.NewFileTypeService(repos.FileTypeRepo),
		RoleService:        roleService,
		PersonService:      services.NewPersonService(repos.PersonRepo),
		AddressService:     services.NewAddressService(repos.AddressRepo),
		MedicineService:    services.NewMedicineService(repos.MedicineRepo),
		ProviderService:    services.NewProviderService(repos.ProviderRepo),
		PaymentService:     externalService.NewPaymentService(),
		ExpenseService:     services.NewExpenseService(repos.ExpenseRepo),
		ImposedService:     services.NewImposedService(repos.ImposedRepo),
		OrderDetailService: services.NewOrderDetailService(repos.OrderDetailRepo),
		JwtTokenService:    jwtTokenService,
		UserService:        services.NewUserService(repos.UserRepo, jwtTokenService, roleService),
		CategoryService:    services.NewCategoryService(repos.CategoryRepo),
	}
}
