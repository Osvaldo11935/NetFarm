package extensions

import (
	"NetFarm/shared/models/common"
	"gorm.io/gorm"
	"math"
)

func Paginate(value interface{}, pagination *common.Pagination, db *gorm.DB) func(db *gorm.DB) *gorm.DB {
	var totalRows int64
	db.Model(value).Count(&totalRows)

	pagination.Total = totalRows
	totalPages := int(math.Ceil(float64(totalRows) / float64(pagination.GetLimit())))
	pagination.TotalPages = totalPages

	return func(db *gorm.DB) *gorm.DB {
		return db.Offset(pagination.GetOffset()).Limit(pagination.GetLimit())
	}
}
