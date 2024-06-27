package setup

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories"
)

type RepositoriesSetup struct {
	OrdersRepo      irepositories.IOrderRepository
	OrderStatusRepo irepositories.IOrderStatusRepository
	OrderFileRepo   irepositories.IOrderFileRepository
	OrderItemRepo   irepositories.IOrderItemRepository
	FileTypeRepo    irepositories.IFileTypeRepository
	RoleRepo        irepositories.IRoleRepository
	PersonRepo      irepositories.IPersonRepository
	UserRepo        irepositories.IUserRepository
	AddressRepo     irepositories.IAddressRepository
	MedicineRepo    irepositories.IMedicineRepository
	ProviderRepo    irepositories.IProviderRepository
	ImposedRepo     irepositories.IImposedRepository
	ExpenseRepo     irepositories.IExpenseRepository
	OrderDetailRepo irepositories.IOrderDetailRepository
	CategoryRepo    irepositories.ICategoryRepository
}

func NewRepositories() *RepositoriesSetup {
	return &RepositoriesSetup{
		OrdersRepo:      repositories.NewOrderRepository(),
		OrderStatusRepo: repositories.NewOrderStatusRepository(),
		OrderFileRepo:   repositories.NewOrderFileRepository(),
		OrderItemRepo:   repositories.NewOrderItemRepository(),
		FileTypeRepo:    repositories.NewFileTypeRepository(),
		RoleRepo:        repositories.NewRoleRepository(),
		PersonRepo:      repositories.NewPersonRepository(),
		UserRepo:        repositories.NewUserRepository(),
		AddressRepo:     repositories.NewAddressRepository(),
		MedicineRepo:    repositories.NewMedicineRepository(),
		ProviderRepo:    repositories.NewProviderRepository(),
		ImposedRepo:     repositories.NewImposedRepository(),
		ExpenseRepo:     repositories.NewExpenseRepository(),
		OrderDetailRepo: repositories.NewOrderDetailRepository(),
		CategoryRepo:    repositories.NewCategoryRepository(),
	}
}
