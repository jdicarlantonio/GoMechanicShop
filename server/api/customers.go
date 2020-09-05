package api

import (
	"fmt"
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

// the whole customer returned by queries
type WholeCustomer struct {
	Id string `json:"id"`
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

func GetCustomersByLastName(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB()
	util.EnableCors(&w, r)
	if(*r).Method == "OPTIONS" {
		return
	}

	var lastName Customer

	err := json.NewDecoder(r.Body).Decode(&lastName)
	util.Check(err)

	customerQuery := `SELECT * FROM customer WHERE lname = $1`
	rows, err := db.Query(customerQuery, lastName.Lname)
	util.Check(err)
	defer rows.Close()

	var customers []WholeCustomer
	for rows.Next() {
		var currentCustomer WholeCustomer
		err := rows.Scan(
			&currentCustomer.Id,
			&currentCustomer.Fname,
			&currentCustomer.Lname,
			&currentCustomer.Phone,
			&currentCustomer.Address)
		util.Check(err)

		customers = append(customers, currentCustomer)
	}

	if len(customers) < 1 {
		var invalidMessage util.Message
		invalidMessage.Message = "Last name does not exist in database"

		out, err := json.Marshal(invalidMessage)
		util.Check(err)

		fmt.Fprintf(w, string(out))
		return
	}

	out, err := json.Marshal(customers)
	util.Check(err)

	fmt.Fprintf(w, string(out))
	defer db.Close()
}

func RemoveCustomer(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB()
	util.EnableCors(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var customer WholeCustomer
	var invalidMessage util.Message
	
	err := json.NewDecoder(r.Body).Decode(&customer)
	util.Check(err)

	// get all the car vins owned by the customer
	vinQuery := `SELECT car_vin FROM owns WHERE customer_id = $1`
	exists, vinList := util.QueryReturnRows(vinQuery, db, customer.Id)
	if !exists {
		invalidMessage.Message = "ID does not exist in database"

		out, err := json.Marshal(invalidMessage)
		util.Check(err)

		fmt.Fprintf(w, string(out))
		return
	}

	// remove customer
	deleteCustomer := `DELETE FROM customer WHERE id = $1`
	_, err = db.Exec(deleteCustomer, customer.Id)
	util.Check(err)

	// remove all of the customers cars
	deleteCars := `DELETE FROM car WHERE vin = $1`
	for _, vin := range vinList {
		_, err := db.Exec(deleteCars, vin)
		util.Check(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}