package request

import (
	"fmt"
	"net/http"
	// "technician/config"
	"technician/db"
	"technician/srv/sse"
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
	sse.Send(technicians, w, r)
	w.WriteHeader(200)
}

func ServiceRequestSchedule(w http.ResponseWriter, r *http.Request) {
	// request := getServiceRequestBody(w, r)
	
	// // dbase := db.GetPostgresDB()
	// technicians := []config.AvailableTechnicians{
	// 	{
	// 		Username: request.Technician,
	// 	},
	// }
	// err := db.StoreRequest(dbase, request, technicians)
	// if err != nil {
	// 	http.Error(w, "couldn't store technicians", http.StatusBadRequest)
	// }
	// sse.Send(technicians, w, r)
	sse.HandleTextMessage(w, r)
	// w.WriteHeader(200)
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