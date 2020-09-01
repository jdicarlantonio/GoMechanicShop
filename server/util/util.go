package util

import (
	"log"
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

// print errors
func Check(err error) {
	if err != nil {
		log.Println(err)
	}
}