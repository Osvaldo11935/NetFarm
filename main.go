package main

import (
	"NetFarm/infrastructure/storages"
	"NetFarm/shared/extensions"
	"NetFarm/webApi/routers/common"
	"NetFarm/webApi/setup"
	"fmt"
	"net/http"
)

func main() {
	// Setup repositories, servicesSetup, and controllersSetup
	repos := setup.NewRepositories()
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
