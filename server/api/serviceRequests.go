package api

import (
	"fmt"
	"log"
	"time"
	"net/http"
	"encoding/json"

	"GoMechanicShop/storage"
	"GoMechanicShop/util"
)

type ServiceRequest struct {
	CustomerFirstName string `json:"fname"`
	CustomerLastName string `json:"lname"`
	CustomerPhone string `json:"phone"`
	CarVin string `json:"vin"`
	Odometer string `json:"odometer"`
	Complaint string `json:"complaint"`
}

type ClosedRequest struct {
	RequestId string `json:"rid"`
	MechanicId string `json:"mid"`
	Comment string `json:"comment"`
	Bill string `json:"bill"`
}

type OpenServiceRequest struct {
	RequestId int `json:"rid"`
	CustomerId int `json:"customer_id"`
	CarVin string `json:"car_vin"`
	Date string `json:"date"`
	Odometer int `json:"odometer"`
	Complaint string `json:"complain"`
}

func AddServiceRequest(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB();
	util.EnableCors(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var invalidMessage util.Message

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
		invalidMessage.Message = "Customer does not exist in database"
		out, err := json.Marshal(invalidMessage)
		util.Check(err)

		fmt.Fprintf(w, string(out))		
		return
	}

	carQuery := `SELECT vin FROM car WHERE vin = $1`
	exists, _ = util.QueryReturn(carQuery, db, serviceRequest.CarVin)
	if !exists {
		invalidMessage.Message = "Car does not exist in database"
		out, err := json.Marshal(invalidMessage)
		util.Check(err)

		fmt.Fprintf(w, string(out))
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
	util.EnableCors(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}
	
	var invalidMessage util.Message

	var closedRequest ClosedRequest
	err := json.NewDecoder(r.Body).Decode(&closedRequest)
	util.Check(err)
	log.Println(closedRequest)

	ridExists := `SELECT rid FROM service_request WHERE rid = $1`
	exists, _ := util.QueryReturn(ridExists, db, closedRequest.RequestId)
	if !exists {
		invalidMessage.Message = "Request does not exist in database"
		out, err := json.Marshal(invalidMessage)
		util.Check(err)

		fmt.Fprintf(w, string(out))
		return
	}

	midExists := `SELECT id FROM mechanic WHERE id = $1`
	exists, _ = util.QueryReturn(midExists, db, closedRequest.MechanicId)
	if !exists {
		invalidMessage.Message = "Employee does not exist in database"
		out, err := json.Marshal(invalidMessage)
		util.Check(err)

		fmt.Fprintf(w, string(out))
		return
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

type Timeline struct {
	StartDate string `json:"startDate"`
	EndDate string `json:"endDate"`
}

func GetOpenServiceRequests(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB()
	util.EnableCors(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	openServiceRequestsQuery := `
		SELECT * FROM service_request 
		WHERE rid NOT IN (SELECT rid FROM closed_request);`

	rows, err := db.Query(openServiceRequestsQuery)
	util.Check(err)
	defer rows.Close()

	var openServiceRequests []OpenServiceRequest
	for rows.Next() {
		request := OpenServiceRequest{}
		err = rows.Scan(
			&request.RequestId,
			&request.CustomerId,
			&request.CarVin,
			&request.Date,
			&request.Odometer,
			&request.Complaint,
		)
		util.Check(err)
		openServiceRequests = append(openServiceRequests, request)
	}

	if len(openServiceRequests) < 1 {
		var invalidMessage util.Message
		invalidMessage.Message = "No service requests exist in database"

		out, err := json.Marshal(invalidMessage)
		util.Check(err)

		fmt.Fprintf(w, string(out))
		return
	}


	out, err := json.Marshal(openServiceRequests)
	util.Check(err)

	fmt.Fprintf(w, string(out))
	defer db.Close()
}