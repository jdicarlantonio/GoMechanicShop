package api

import (
	"fmt"
	"log"
	"net/http"
	"encoding/json"
	"database/sql"

	"GoMechanicShop/storage"
	"GoMechanicShop/util"
)

// holds information on the car and its owner
type CarOwnerRelation struct {
	Vin string `json:"vin"`
	Make string `json:"make"`
	Model string `json:"model"`
	Year string `json:"year"`
	OwnerFirstName string `json:"fname"`
	OwnerLastName string `json:"lname"`
	OwnerPhone string `json:"phone"`
}

func AddCar(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB();
	util.EnableCors(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var invalidMessage util.Message

	var carOwnerRelation CarOwnerRelation
	err := json.NewDecoder(r.Body).Decode(&carOwnerRelation)
	util.Check(err)

	// check if customer exists
	customerExistsQuery := `SELECT id FROM customer WHERE fname = $1 AND lname = $2 AND phone = $3`
	row := db.QueryRow(customerExistsQuery, carOwnerRelation.OwnerFirstName, carOwnerRelation.OwnerLastName, carOwnerRelation.OwnerPhone)

	var customerId int
	err = row.Scan(&customerId)
	if err == sql.ErrNoRows {
		invalidMessage.Message = "Customer does not exist in database"
		out, err := json.Marshal(invalidMessage)
		util.Check(err)

		fmt.Fprintf(w, string(out))
		return
	} else {
		util.Check(err)
	}

	log.Println(carOwnerRelation)

	carUpdate := `INSERT INTO car (vin, make, model, year) VALUES ($1, $2, $3, $4)`
	_, err = db.Exec(carUpdate, carOwnerRelation.Vin, carOwnerRelation.Make, carOwnerRelation.Model, carOwnerRelation.Year)
	util.Check(err)

	// create the owner relation
	ownerUpdate := `INSERT INTO owns (customer_id, car_vin) VALUES ($1, $2)`
	_, err = db.Exec(ownerUpdate, customerId, carOwnerRelation.Vin)
	util.Check(err)

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
