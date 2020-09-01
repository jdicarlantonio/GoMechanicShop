package api

import (
	"log"
	"net/http"
	"encoding/json"

	"GoMechanicShop/storage"
	"GoMechanicShop/util"
)

type Mechanic struct {
	Fname string `json:"fname"`
	Lname string `json:"lname"`
	Experience string `json:"experience"`
}

func AddMechanic(w http.ResponseWriter, r *http.Request) {
	db := storage.ConnectToDB()
	util.EnableCors(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	var mechanic Mechanic
	
	err := json.NewDecoder(r.Body).Decode(&mechanic)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println(mechanic)

	mechanicUpdate := `INSERT INTO mechanic (fname, lname, experience) VALUES ($1, $2, $3)`
	_, err = db.Exec(mechanicUpdate, mechanic.Fname, mechanic.Lname, mechanic.Experience)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		panic(err)
	}

	w.WriteHeader(http.StatusOK)
	defer db.Close()
}