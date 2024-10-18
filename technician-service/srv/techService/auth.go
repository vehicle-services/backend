package techService

import (
	"fmt"
	"net/http"
	"technician/config"
	"technician/db"
)

var jwtSecret = []byte("your-secret-key")

func Signup(w http.ResponseWriter, r *http.Request) {
	technician := getBody(w, r)
	
	dbase := db.GetPostgresDB()
	err := db.CreateTechnician(dbase, technician)
	if err != nil {
		http.Error(w, "can not create technician", http.StatusBadRequest)
		fmt.Println(err)
		return
	}

	writeResponse(technician, w)
}

func Login(w http.ResponseWriter, r *http.Request) {
	technician := getBody(w, r)

	token, err := GenerateToken(technician.Username, technician.Password)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return
	}
	auth := config.Auth{
		Token: token,
	}
	writeResponseToken(auth, w)
}