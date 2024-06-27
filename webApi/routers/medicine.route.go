package routers

import (
	"NetFarm/webApi/controllers"
	"github.com/gin-gonic/gin"
)

func AddMedicineRoute(route *gin.RouterGroup, controller *controllers.MedicineController) {
	medicineRouter := route.Group("/medicines")
	medicineRouter.GET("", controller.HandleFindAllMedicine)
	medicineRouter.GET("provider/:providerId", controller.HandleFindMedicineByProviderId)
	medicineRouter.POST("provider/:providerId", controller.HandleCreate)
	medicineRouter.POST(":medicineId", controller.HandleCreateMedicineCategory)
	medicineRouter.GET(":medicineId", controller.HandleFindMedicineById)
	medicineRouter.PUT(":medicineId", controller.HandleUpdateMedicine)
	medicineRouter.DELETE(":medicineId", controller.HandleDeleteMedicine)

}
