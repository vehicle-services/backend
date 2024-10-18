package techService

import (
	"net/http"
	"technician/db"
)

func GetTechnician(w http.ResponseWriter, r *http.Request) {
	username := protectedHandler(w, r)
	technician := getBody(w, r)
	if (username != technician.Username){
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}
	
	dbase := db.GetPostgresDB()
	technician, err := db.GetTechnician(dbase, technician)
	if err != nil {
		http.Error(w, "can not get technician", http.StatusBadRequest)
	}

	writeResponse(technician, w)
}

func UpdateShopDetails(w http.ResponseWriter, r *http.Request) {
	username := protectedHandler(w, r)
	technician := getBody(w, r)
	if (username != technician.Username){
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}
	
	dbase := db.GetPostgresDB()
	technician, err := db.UpdateShopDetails(dbase, technician)
	if err != nil {
		http.Error(w, "can not update technician", http.StatusBadRequest)
	}

	writeResponse(technician, w)
}

func OffboardTechnician(w http.ResponseWriter, r *http.Request) {
	username := protectedHandler(w, r)
	technician := getBody(w, r)
	if (username != technician.Username){
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}
	
	dbase := db.GetPostgresDB()
	technician, err := db.DeleteTechnician(dbase, technician)
	if err != nil {
		http.Error(w, "can not get technician", http.StatusBadRequest)
	}

	writeResponse(technician, w)
}
