package techService

import (
	"encoding/json"
	"io"
	"net/http"
	"technician/config"
	"log"
	"github.com/golang-jwt/jwt/v4"
	"fmt"
	"time"
)

func getBody(w http.ResponseWriter, r *http.Request) config.Technician {
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

func writeResponse(technician config.Technician, w http.ResponseWriter) {
	response, err := json.Marshal(technician)
	if err != nil {
		http.Error(w, "can not marshal response", http.StatusBadGateway)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

func writeResponseToken(token config.Auth, w http.ResponseWriter) {
	response, err := json.Marshal(token)
	if err != nil {
		http.Error(w, "can not marshal response", http.StatusBadGateway)
		return
	}
	w.WriteHeader(200)
	w.Write(response)
}

func GenerateToken(username, password string) (string, error) {
	if username == "" || password == "" {
		return "", fmt.Errorf("username and password are required")
	}

	claims := config.Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)), // Token expires in 1 hour
			Issuer:    "mechnanix",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", fmt.Errorf("error signing the token: %v", err)
	}

	return tokenString, nil
}

func ValidateToken(tokenString string) (*config.Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &config.Claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*config.Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func protectedHandler(w http.ResponseWriter, r *http.Request) string {
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		http.Error(w, "Authorization header missing", http.StatusUnauthorized)
		return ""
	}
	claims, err := ValidateToken(tokenString)
	if err != nil {
		http.Error(w, err.Error(), http.StatusUnauthorized)
		return ""
	}
	return claims.Username
}