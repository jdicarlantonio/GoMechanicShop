package api

import (
	"log"
	"net/http"
	"encoding/json"

	"GoMechanicShop/storage"
	"GoMechanicShop/util"
)

type Customer struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Phone string `json:"phone"`
	Address string `json:"address"`
}

func AddCustomer(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB()
	util.EnableCors(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var customer Customer
	
	err := json.NewDecoder(r.Body).Decode(&customer)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(customer)

	customerUpdate := `INSERT INTO customer (fname, lname, phone, address) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(customerUpdate, customer.Fname, customer.Lname, customer.Phone, customer.Address)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}