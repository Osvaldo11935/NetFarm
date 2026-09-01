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

	sqlDB, err := db.DB()
	if err != nil {
		return nil, err
	}

	sqlDB.SetMaxOpenConns(5)
	sqlDB.SetMaxIdleConns(2)

	err = db.AutoMigrate(
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

	if err != nil {
		return nil, err
	}

	entities.InitialValueRole(db)
	entities.InitialValueUser(db)
	entities.InitialValueStatus(db)

	return db, nil
}