package entities

type MedicineCategory struct {
	CategoryId string `gorm:"column:CategoryId;primaryKey" json:"categoryId"`
	MedicineId string `gorm:"column:MedicineId;primaryKey" json:"medicineId"`
}
