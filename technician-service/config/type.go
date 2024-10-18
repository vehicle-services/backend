package config

import (
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type Technician struct {
	TechnicianID string  `json:"technicianid"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	ShopName     string  `json:"shop_name"`
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"long"`
	MobileNumber int64   `json:"mobile_no"`
	Description  string  `json:"shop_desc"`
	Open         string  `json:"currently_open"`
	Token        string  `json:"token"`
	OpenTime     string  `json:"open_time"`
	CloseTime    string  `json:"close_time"`
}

type ServiceRequest struct {
	RequestID           string    `json:"request_id"`
	UserID              string    `json:"user_id"`
	ServiceDescription  string    `json:"issue_description"`
	VehicleType         string    `json:"vehicle_type"`
	Model_name          string    `json:"model_name"`
	ServiceType         string    `json:"service_type"`
	Date                time.Time `json:"date"`
	Latitude            float64   `json:"lat"`
	Longitude           float64   `json:"long"`
	Distance            int64     `json:"distance"`
	NumberOfTechnicians int       `json:"num_tech"`
	Status              string    `json:"status"`
	Technician          string    `json:"technician"`
}

type AvailableTechnicians struct {
	Username     string  `json:"username"`
	ShopName     string  `json:"shop_name"`
	Latitude     float64 `json:"lat"`
	Longitude    float64 `json:"long"`
	MobileNumber int64   `json:"mobile_no"`
	Description  string  `json:"shop_desc"`
	Open         string  `json:"shop_open"`
	Distance     float64 `json:"distance"`
	OpenTime     string  `json:"open_time"`
	CloseTime    string  `json:"close_time"`
}

type TechnicianRequest struct {
	Username string         `json:"username"`
	Requests ServiceRequest `json:"service_request"`
}

type Auth struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
