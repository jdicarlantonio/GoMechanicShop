package api

import (
	"log"
	"time"
	"net/http"
	"encoding/json"
//	"database/sql"

	"GoMechanicShop/storage"
	"GoMechanicShop/util"
)

type ServiceRequest struct {
	CustomerFirstName string `json:"fname"`
	CustomerLastName string `json:"lname"`
	CustomerPhone string `json:"phone"`
	CarVin string `json:"vin"`
	Odometer int `json:"odometer"`
	Complaint string `json:"complaint"`
}

type ClosedRequest struct {
	RequestId int `json:"rid"`
	MechanicId int `json:"mid"`
	Comment string `json:"comment"`
	Bill int `json:"bill"`
}

func AddServiceRequest(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB();

	var serviceRequest ServiceRequest
	err := json.NewDecoder(r.Body).Decode(&serviceRequest)
	util.Check(err)
	log.Println(serviceRequest)

	customerQuery := `SELECT id FROM customer WHERE fname = $1 AND lname = $2 AND phone = $3`
	exists, customerId := util.QueryReturn(
		customerQuery,
		db,
		serviceRequest.CustomerFirstName, 
		serviceRequest.CustomerLastName, 
		serviceRequest.CustomerPhone)
	if !exists {
		log.Println("create customer")
		return
	}

	carQuery := `SELECT vin FROM car WHERE vin = $1`
	exists, _ = util.QueryReturn(carQuery, db, serviceRequest.CarVin)
	if !exists {
		log.Println("create car")
		return
	}

	serviceRequestUpdate := `
		INSERT INTO service_request (customer_id, car_vin, date, odometer, complain) 
		VALUES ($1, $2, $3, $4, $5)`

	dt := time.Now()
	_, err = db.Exec(
		serviceRequestUpdate, 
		customerId, 
		serviceRequest.CarVin, 
		dt.Format("01-02-2006"), 
		serviceRequest.Odometer, 
		serviceRequest.Complaint)
	util.Check(err)

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}

func CloseServiceRequest(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB()

	var closedRequest ClosedRequest
	err := json.NewDecoder(r.Body).Decode(&closedRequest)
	util.Check(err)
	log.Println(closedRequest)

	ridExists := `SELECT rid FROM service_request WHERE rid = $1`
	exists, _ := util.QueryReturn(ridExists, db, closedRequest.RequestId)
	if !exists {
		log.Println("request doesn't exists")
	}

	midExists := `SELECT id FROM mechanic WHERE id = $1`
	exists, _ = util.QueryReturn(midExists, db, closedRequest.MechanicId)
	if !exists {
		log.Println("mechanic doesn't exists")
	}

	closeServiceRequest := `
		INSERT INTO closed_request (rid, mid, date, comment, bill)
		VALUES ($1, $2, $3, $4, $5)`

	dt := time.Now()
	_, err = db.Exec(
		closeServiceRequest,
		closedRequest.RequestId,
		closedRequest.MechanicId,
		dt.Format("01-02-2006"), 
		closedRequest.Comment,
		closedRequest.Bill)

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}
