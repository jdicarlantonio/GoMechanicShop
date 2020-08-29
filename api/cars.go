package api

import (
	"log"
	"net/http"
	"encoding/json"
	"database/sql"

	"GoMechanicShop/storage"
)

// holds information on the car and its owner
type CarOwnerRelation struct {
	Vin string `json:"vin"`
	Make string `json:"make"`
	Model string `json:"model"`
	Year int `json:"year"`
	OwnerFirstName string `json:"fname"`
	OwnerLastName string `json:"lname"`
	OwnerPhone string `json:"phone"`
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB();

	var carOwnerRelation CarOwnerRelation
	err := json.NewDecoder(r.Body).Decode(&carOwnerRelation)
	Check(err, w)
	log.Println(carOwnerRelation)

	// check if customer exists
	customerExistsQuery := `SELECT id FROM customer WHERE fname = $1 AND lname = $2 AND phone = $3`
	row := db.QueryRow(customerExistsQuery, carOwnerRelation.OwnerFirstName, carOwnerRelation.OwnerLastName, carOwnerRelation.OwnerPhone)

	var customerId int
	err = row.Scan(&customerId)
	if err == sql.ErrNoRows {
		log.Print("create customer")
	} else {
		Check(err, w)
	}

	carUpdate := `INSERT INTO car (vin, make, model, year) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(carUpdate, carOwnerRelation.Vin, carOwnerRelation.Make, carOwnerRelation.Model, carOwnerRelation.Year)
	Check(err, w)

	// create the owner relation
	ownerUpdate := `INSERT INTO owns (customer_id, car_vin) VALUES ($1, $2)`
	_, err = db.Exec(ownerUpdate, customerId, carOwnerRelation.Vin)
	Check(err, w)

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func Check(err error, w http.ResponseWriter) {
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}
}