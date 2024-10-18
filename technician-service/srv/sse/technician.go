package sse

import (
	"encoding/json"
	"fmt"
	"net/http"
	"technician/config"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow all origins for testing purposes
	},
}

func HandleTextMessage(w http.ResponseWriter, r *http.Request) {
	// Process the incoming message
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	
	// Send a response
	response := "Hello from server!"
	err = conn.WriteMessage(websocket.TextMessage, []byte(response))
	if err != nil {
		fmt.Println(err)
	}
}

func Send(technicians []config.AvailableTechnicians, w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")
	fmt.Println("1----------------------")

	if flusher, ok := w.(http.Flusher); ok {
		flusher.Flush()
	}
	fmt.Println("2----------------------")

	for i := 0; i < len(technicians); i++ {
		event, err := json.Marshal(technicians[i])
		if err != nil {
			fmt.Println("Error marshalling JSON:", err)
			return
		}
		fmt.Println("3----------------------")

		select {
		case <-r.Context().Done():
			fmt.Println("Client disconnected")
			return
		default:
			fmt.Println("4----------------------")
			fmt.Fprintf(w, "data: %s\n\n", event)
			if flusher, ok := w.(http.Flusher); ok {
				flusher.Flush()
			}
		}
	}
}
