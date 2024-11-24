package config

import (
	"github.com/golang-jwt/jwt/v4"
)

type Technician struct {
	TechnicianID string  `json:"technicianid"`
	Username     string  `json:"username"`
	Password     string  `json:"password"`
	ShopName     string  `json:"shopName"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	MobileNumber int64   `json:"mobileNumber"`
	Description  string  `json:"shopDesc"`
	Open         string  `json:"currentlyOpen"`
	Token        string  `json:"token"`
	OpenTime     string  `json:"openTime"`
	CloseTime    string  `json:"closeTime"`
}

type ServiceRequest struct {
	RequestID           string  `json:"serviceRequestUUID"`
	UserID              string  `json:"userEmail"`
	ServiceDescription  string  `json:"serviceDescription"`
	VehicleType         string  `json:"vehicleType"`
	Model_name          string  `json:"modelName"`
	ServiceType         string  `json:"serviceType"`
	Date                string  `json:"date"`
	Time                string  `json:"time"`
	Latitude            float64 `json:"latitude"`
	Longitude           float64 `json:"longitude"`
	Distance            int64   `json:"distance"`
	NumberOfTechnicians int     `json:"numTech"`
	Status              string  `json:"status"`
	Technician          string  `json:"technician"`
	RequestType         string  `json:"requestType"`
}

type AvailableTechnicians struct {
	Username     string  `json:"userName"`
	ShopName     string  `json:"shopName"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	MobileNumber int64   `json:"mobileNo"`
	Description  string  `json:"shopDesc"`
	Open         string  `json:"shop_open"`
	Distance     float64 `json:"distance"`
	OpenTime     string  `json:"openTime"`
	CloseTime    string  `json:"closeTime"`
}

type TechnicianRequest struct {
	Username string         `json:"username"`
	Requests ServiceRequest `json:"serviceRequest"`
}

type Auth struct {
	Token string `json:"token"`
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}
