package api

import (
	"log"
	"net/http"
	"encoding/json"
//	"database/sql"

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
	util.Check(err)
	log.Println(customer)

	customerUpdate := `INSERT INTO customer (fname, lname, phone, address) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(customerUpdate, customer.Fname, customer.Lname, customer.Phone, customer.Address)
	util.Check(err)

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func RemoveCustomer(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB()
	util.EnableCors(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var customer Customer
	
	err := json.NewDecoder(r.Body).Decode(&customer)
	util.Check(err)
	log.Println(customer)

	// get the ID of the customer
	customerIDQuery := `
		SELECT id FROM customer 
		WHERE fname = $1 AND lname = $2 AND phone = $3`
	exists, customerID := util.QueryReturn(customerIDQuery, db, customer.Fname, customer.Lname, customer.Phone)
	if !exists {
		log.Println("Customer does not exist")
		w.WriteHeader(http.StatusNoContent)
		return
	}

	// need to remove customer's cars and service requests
	// since the provided schema does not cascade delete, we have to do a lot of this manually
	
	// first get vin of all customer's cars
	ownerVinQuery := `
		SELECT car_vin FROM owns 
		WHERE customer_id = $1`

	exists, ownerVins := util.QueryReturnRows(ownerVinQuery, db, customerID)
	if !exists {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	log.Println(ownerVins)

	// now need the rid of all the customer's service requests
	serviceRequestIDQuery := `
	SELECT rid FROM service_request
	WHERE customer_id = $1`

	exists, ownerRids := util.QueryReturnRows(serviceRequestIDQuery, db, customerID)
	if !exists {
		w.WriteHeader(http.StatusNoContent)
		return
	}
	log.Println(ownerRids)

	
	deleteOwnsRelation := `DELETE FROM owns WHERE customer_id = $1`
	_, err = db.Exec(deleteOwnsRelation, customerID)
	util.Check(err)

	deleteClosedServiceRequest := `DELETE FROM closed_request WHERE rid = $1`
	for _, rid := range ownerRids {
		_, err = db.Exec(deleteClosedServiceRequest, rid)
		util.Check(err)
	}
	
	deleteServiceRequest := `DELETE FROM service_request WHERE customer_id = $1`
	_, err = db.Exec(deleteServiceRequest, customerID)
	util.Check(err)

	deleteCars := `DELETE FROM car WHERE vin = $1`
	for _, vin := range ownerVins {
		_, err = db.Exec(deleteCars, vin)
		util.Check(err)
	}
	
	w.WriteHeader(http.StatusOK)
	defer db.Close()
}