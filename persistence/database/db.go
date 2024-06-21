package database

import (
	"NetFarm/domain/entities"
	"NetFarm/shared/extensions"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() (*gorm.DB, error) {

	connectionString := extensions.GetEnv("POSTGRES_CONNECTION_STRING")

	db, err := gorm.Open(postgres.Open(connectionString), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	errMigration := db.AutoMigrate(
		&entities.Role{},
		&entities.User{},
		&entities.Address{},
		&entities.OrderStatus{},
		&entities.Order{},
		&entities.FileType{},
		&entities.OrderFile{},
		&entities.OrderItem{},
		&entities.Person{},
		&entities.Provider{},
		&entities.Medicine{},
		&entities.OrderDetail{},
		&entities.Expense{},
		&entities.Imposed{},
		&entities.Category{},
		&entities.MedicineCategory{},
	)

	entities.InitialValueRole(db)
	entities.InitialValueUser(db)
	entities.InitialValueStatus(db)

	if errMigration != nil {
		return nil, errMigration
	}

	return db, nil
}
