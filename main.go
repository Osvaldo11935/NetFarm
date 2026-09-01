package main

import (
	"NetFarm/infrastructure/storages"
	"NetFarm/shared/extensions"
	"NetFarm/webApi/routers/common"
	"NetFarm/webApi/setup"
	"fmt"
	"net/http"
	"NetFarm/persistence/database"
)

func main() {
	db, errDb := database.NewDatabase()
	if errDb != nil {
		fmt.Println("Erro ao conectar ao banco:", errDb)
		return
	}

	repos := setup.NewRepositories(db)

	servicesSetup := setup.NewServices(repos)
	storage := storages.NewGoogleStorage()
	controllersSetup := setup.NewControllers(servicesSetup, &storage)

	// Setup router
	routes := common.NewRouter(
		controllersSetup.OrderController,
		controllersSetup.OrderFileController,
		controllersSetup.OrderStatusController,
		controllersSetup.OrderItemController,
		controllersSetup.AddressController,
		controllersSetup.FileTypeController,
		controllersSetup.PersonController,
		controllersSetup.RoleController,
		controllersSetup.UserController,
		controllersSetup.MedicineController,
		controllersSetup.ProviderController,
		controllersSetup.PaymentController,
		controllersSetup.ExpenseController,
		controllersSetup.ImposedController,
		controllersSetup.OrderDetailController,
		controllersSetup.CategoryController,
	)

	port := extensions.GetEnv("PORT")
	
	server := &http.Server{
		Addr:    ":"+port,
		Handler: routes,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err)
	}
}
