package main

import (
	"log"
	"net/http"

    "GoMechanicShop/api"
)

func main() {
	http.HandleFunc("/addCustomer", api.AddCustomer)
	http.HandleFunc("/addMechanic", api.AddMechanic)
	http.HandleFunc("/addCar", api.AddCar)
	http.HandleFunc("/addServiceRequest", api.AddServiceRequest)
	http.HandleFunc("/closeServiceRequest", api.CloseServiceRequest)
	http.HandleFunc("/getOpenServiceRequests", api.GetOpenServiceRequests)
	http.HandleFunc("/removeCustomer", api.RemoveCustomer)
	http.HandleFunc("/getCustomersByLastName", api.GetCustomersByLastName)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
