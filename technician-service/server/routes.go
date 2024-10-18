package server

import (
	"net/http"
	"technician/srv/request"
	// "technician/srv/sse"
	"technician/srv/techService"

	"github.com/gorilla/mux"
)

type Handler struct {
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/version", h.version).Methods("GET")
	router.HandleFunc("/technician", h.GetTechnician).Methods("GET")
	router.HandleFunc("/signup", h.Signup).Methods("POST")
	router.HandleFunc("/login", h.Login).Methods("POST")
	router.HandleFunc("/shop", h.UpdateShopDetails).Methods("PUT")
	router.HandleFunc("/technician", h.OffboardTechnician).Methods("DELETE")
	router.HandleFunc("/availabletechnicians", h.AvailableTechnicians).Methods("POST")
	router.HandleFunc("/serviceRequest", h.ServiceRequest).Methods("POST")
	router.HandleFunc("/scheduleRequest", h.ServiceRequestSchedule).Methods("POST")
	router.HandleFunc("/requests", h.GetServiceRequests).Methods("POST")
	router.HandleFunc("/requestStatus", h.UpdateRequestStatus).Methods("PATCH")
}

func (h *Handler) version(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain")
	w.Write([]byte("1.0"))
}

func (h *Handler) GetTechnician(w http.ResponseWriter, r *http.Request) {
	techService.GetTechnician(w, r)
}

func (h *Handler) Signup(w http.ResponseWriter, r *http.Request) {
	techService.Signup(w, r)
}

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	techService.Login(w, r)
}

func (h *Handler) UpdateShopDetails(w http.ResponseWriter, r *http.Request) {
	techService.UpdateShopDetails(w, r)
}

func (h *Handler) OffboardTechnician(w http.ResponseWriter, r *http.Request) {
	techService.OffboardTechnician(w, r)
}

func (h *Handler) AvailableTechnicians(w http.ResponseWriter, r *http.Request) {
	request.AvailableTechnicians(w, r)
}

func (h *Handler) ServiceRequest(w http.ResponseWriter, r *http.Request) {
	request.ServiceRequest(w, r)
}

func (h *Handler) ServiceRequestSchedule(w http.ResponseWriter, r *http.Request) {
	request.ServiceRequestSchedule(w, r)
}

func (h *Handler) GetServiceRequests(w http.ResponseWriter, r *http.Request) {
	request.GetServiceRequests(w, r)
}

func (h *Handler) UpdateRequestStatus(w http.ResponseWriter, r *http.Request) {
	request.UpdateRequestStatus(w, r)
}
