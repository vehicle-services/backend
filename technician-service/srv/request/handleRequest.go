package request

import (
	"fmt"
	"net/http"
	"technician/config"
	"technician/db"
)

func ServiceRequest(w http.ResponseWriter, r *http.Request) {
	request := getServiceRequestBody(w, r)
	
	dbase := db.GetPostgresDB()
	technicians, err := db.GetTechniciansLocation(dbase, request)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "can't get technicians", http.StatusBadRequest)
		return
	}
	err = db.StoreRequest(dbase, request, technicians)
	if err != nil {
		fmt.Println(err)
		http.Error(w, "couldn't store request", http.StatusBadRequest)
	}
	w.WriteHeader(200)
}

func ServiceRequestSchedule(w http.ResponseWriter, r *http.Request) {
	request := getServiceRequestBody(w, r)
	
	dbase := db.GetPostgresDB()
	technicians := []config.AvailableTechnicians{
		{
			Username: request.Technician,
		},
	}
	err := db.StoreRequest(dbase, request, technicians)
	if err != nil {
		http.Error(w, "couldn't store technicians", http.StatusBadRequest)
	}
	w.WriteHeader(200)
}

func AvailableTechnicians(w http.ResponseWriter, r *http.Request) {
	request := getServiceRequestBody(w, r)
	
	dbase := db.GetPostgresDB()
	technicians, err := db.GetTechniciansLocation(dbase, request)
	if err != nil {
		http.Error(w, "can not get technicians", http.StatusBadRequest)
	}
	writeAvailableTechnicianResponse(technicians, w)
}

func DeleteRequest(w http.ResponseWriter, r *http.Request){
	request := getServiceRequestBody(w, r)
	
	dbase := db.GetPostgresDB()
	flag := true
	if request.Technician != "" {
		flag = false
	}
	err := db.RemoveRequest(dbase, request, flag)
	if err != nil {
		http.Error(w, "can not delete request", http.StatusBadRequest)
	}
	writeResponseRequest(request, w)
}