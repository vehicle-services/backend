package request

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"technician/db"
)

func GetServiceRequests(w http.ResponseWriter, r *http.Request) {
	endpoint := "https://"

	getData := getRequestBody(w, r)
	body, err := json.Marshal(getData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPost, endpoint, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create request: %v", err), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to make request: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	request := getBody(w, resp.Body)
	writeResponse(request, w)
}

func UpdateRequestStatus(w http.ResponseWriter, r *http.Request) {
	endpoint := "https://"

	patchData := getRequestStatusBody(w, r)
	body, err := json.Marshal(patchData)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to marshal request body: %v", err), http.StatusInternalServerError)
		return
	}

	req, err := http.NewRequest(http.MethodPatch, endpoint, bytes.NewBuffer(body))
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to create request: %v", err), http.StatusInternalServerError)
		return
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Error(w, fmt.Sprintf("Failed to make request: %v", err), http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	dbase := db.GetPostgresDB()

	removeFlag := patchData.Status == "A"

	err = db.RemoveRequest(dbase, patchData, removeFlag)
	if err != nil {
		http.Error(w, "couldn't remove request", http.StatusBadRequest)
	}

	request := getBody(w, resp.Body)
	writeResponse(request, w)
}
