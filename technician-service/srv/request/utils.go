package request

import (
	"log"
	"io"
	"encoding/json"
	"technician/config"
	"net/http"
)

func getBody(w http.ResponseWriter, r io.ReadCloser) config.TechnicianRequest {
	body, err := io.ReadAll(r)
	if err != nil {
		log.Printf("can not read body")
		http.Error(w, "can't read body", http.StatusBadRequest)
	}
	var request config.TechnicianRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "can not unmarshall body", http.StatusBadRequest)
	}
	return request
}

func getRequestBody(w http.ResponseWriter, r *http.Request) config.Technician {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("can not read body")
		http.Error(w, "can't read body", http.StatusBadRequest)
	}
	var technician config.Technician
	err = json.Unmarshal(body, &technician)
	if err != nil {
		http.Error(w, "can not unmarshall body", http.StatusBadRequest)
	}
	return technician
}

func getServiceRequestBody(w http.ResponseWriter, r *http.Request) config.ServiceRequest {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("can not read body")
		http.Error(w, "can't read body", http.StatusBadRequest)
	}
	var request config.ServiceRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "can not unmarshall body", http.StatusBadRequest)
	}
	return request
}

func getRequestStatusBody(w http.ResponseWriter, r *http.Request) config.ServiceRequest {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("can not read body")
		http.Error(w, "can't read body", http.StatusBadRequest)
	}
	var request config.ServiceRequest
	err = json.Unmarshal(body, &request)
	if err != nil {
		http.Error(w, "can not unmarshall body", http.StatusBadRequest)
	}
	return request
}

func writeResponse(request config.TechnicianRequest, w http.ResponseWriter) {
	response, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "can not marshal response", http.StatusBadGateway)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

func writeResponseRequest(request config.ServiceRequest, w http.ResponseWriter) {
	response, err := json.Marshal(request)
	if err != nil {
		http.Error(w, "can not marshal response", http.StatusBadGateway)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

func writeAvailableTechnicianResponse(technicians []config.AvailableTechnicians, w http.ResponseWriter) {
	response, err := json.Marshal(technicians)
	if err != nil {
		http.Error(w, "can not marshal response", http.StatusBadGateway)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}