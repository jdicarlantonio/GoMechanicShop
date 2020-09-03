package util

import (
	"log"
	"net/http"
	"database/sql"
)

// checks if the query returned a row, returns true if so, false if not
// returns the row of the query if it exists as a string
// only use this if you're expecting only one row to be returned by the query
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

// same as above but returns multiple rows
func QueryReturnRows(query string, db *sql.DB, arg ...interface{}) (bool, []string) {
	rows, err := db.Query(query, arg...)
	if err == sql.ErrNoRows {
		return false, []string{}
	} else {
		Check(err)
	}

	var items []string
	for rows.Next() {
		var currentItem string
		err := rows.Scan(&currentItem)
		if err != nil {
			Check(err)
		}

		items = append(items, currentItem)
	}

	return true, items
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