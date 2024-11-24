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
	router.HandleFunc("/previousRequests", h.GetServiceRequests).Methods("POST")
	router.HandleFunc("/currentRequests", h.GetActiveServiceRequests).Methods("POST")
	router.HandleFunc("/requestStatus", h.UpdateRequestStatus).Methods("PATCH")
	router.HandleFunc("/deleteRequest", h.DeleteRequest).Methods("POST")
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
	techService.GetServiceRequests(w, r)
}

func (h *Handler) GetActiveServiceRequests(w http.ResponseWriter, r *http.Request) {
	techService.GetActiveServiceRequests(w, r)
}

func (h *Handler) UpdateRequestStatus(w http.ResponseWriter, r *http.Request) {
	request.UpdateRequestStatus(w, r)
}

func (h *Handler) DeleteRequest(w http.ResponseWriter, r *http.Request) {
	request.DeleteRequest(w, r)
}