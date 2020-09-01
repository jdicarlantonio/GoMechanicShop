package util

import (
	"log"
	"net/http"
	"database/sql"
)

// checks if the query returned a row, returns true if so, false if not
// returns the row of the query if it exists as a string
func QueryReturn(query string, db *sql.DB, arg ...interface{}) (bool, string) {
	row := db.QueryRow(query, arg...)

	var rowReturned string
	err := row.Scan(&rowReturned)
	if err == sql.ErrNoRows {
		return false, ""
	} else {
		Check(err)
		return true, rowReturned
	}
}

// quick and dirty CORS enable
// this is only for the testing stage and should be handled differently in production
func EnableCors(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
    (*w).Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
    (*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}

// print errors
func Check(err error) {
	if err != nil {
		log.Println(err)
	}
}