package techService

import (
	"net/http"
	"technician/db"
)

func GetServiceRequests(w http.ResponseWriter, r *http.Request){
	username := protectedHandler(w, r)
	technician := getBody(w, r)
	if (username != technician.Username){
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}
	dbase := db.GetPostgresDB()
	requests, err := db.GetRequests(dbase, technician)
	if err != nil {
		http.Error(w, "can not get requests", http.StatusBadRequest)
	}

	writeResponseRequest(requests, w)
}

func GetActiveServiceRequests(w http.ResponseWriter, r *http.Request){
	username := protectedHandler(w, r)
	technician := getBody(w, r)
	if (username != technician.Username){
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return
	}
	dbase := db.GetPostgresDB()
	requests, err := db.GetActiveRequests(dbase, technician)
	if err != nil {
		http.Error(w, "can not get requests", http.StatusBadRequest)
	}

	writeResponseRequest(requests, w)
}

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
