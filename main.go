package main

import (
	"fmt"
	"log"
	"net/http"

    "GoMechanicShop/api"
)

func testPage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello DB")
}

func main() {
    http.HandleFunc("/", testPage)
	http.HandleFunc("/addCustomer", api.AddCustomer)
	http.HandleFunc("/addMechanic", api.AddMechanic)
	http.HandleFunc("/addCar", api.AddCar)

	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}
