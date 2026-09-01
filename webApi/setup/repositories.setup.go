package setup

import (
	"NetFarm/application/interfaces/irepositories"
	"NetFarm/persistence/repositories"
	"gorm.io/gorm"
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

func NewRepositories(db *gorm.DB) *RepositoriesSetup {
	return &RepositoriesSetup{
		OrdersRepo:      repositories.NewOrderRepository(db),
		OrderStatusRepo: repositories.NewOrderStatusRepository(db),
		OrderFileRepo:   repositories.NewOrderFileRepository(db),
		OrderItemRepo:   repositories.NewOrderItemRepository(db),
		FileTypeRepo:    repositories.NewFileTypeRepository(db),
		RoleRepo:        repositories.NewRoleRepository(db),
		PersonRepo:      repositories.NewPersonRepository(db),
		UserRepo:        repositories.NewUserRepository(db),
		AddressRepo:     repositories.NewAddressRepository(db),
		MedicineRepo:    repositories.NewMedicineRepository(db),
		ProviderRepo:    repositories.NewProviderRepository(db),
		ImposedRepo:     repositories.NewImposedRepository(db),
		ExpenseRepo:     repositories.NewExpenseRepository(db),
		OrderDetailRepo: repositories.NewOrderDetailRepository(db),
		CategoryRepo:    repositories.NewCategoryRepository(db),
	}
}
