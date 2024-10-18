package db

import (
	"database/sql"
	"fmt"
	"technician/config"

	"github.com/google/uuid"
)

func CreateTechnician(db *sql.DB, technician config.Technician) error {
	var existingUsername string
	err := db.QueryRow("SELECT username FROM technician WHERE username = $1", technician.Username).Scan(&existingUsername)
	if err != nil && err != sql.ErrNoRows {
		return fmt.Errorf("error checking for existing username: %v", err)
	}
	if existingUsername != "" {
		return fmt.Errorf("username already exists")
	}

	technicianID := uuid.New().String()

	_, err = db.Exec(
		"INSERT INTO technician (technicianid, username, password, shop_name, mobile_no, shop_desc, location, currently_open, open_time, close_time) VALUES ($1, $2, $3, $4, $5, $6, ST_SetSRID(ST_MakePoint($7, $8), 4326), $9, $10, $11)",
		technicianID, technician.Username, technician.Password, technician.ShopName, technician.MobileNumber, technician.Description, technician.Longitude, technician.Latitude, technician.Open, technician.OpenTime, technician.CloseTime,
	)
	if err != nil {
		return fmt.Errorf("error inserting technician %v", err)
	}

	return nil
}

func GetTechnician(db *sql.DB, technician config.Technician) (config.Technician, error) {
	row := db.QueryRow("SELECT technicianid, username, shop_name, mobile_no, shop_desc, currently_open, ST_X(location::geometry) AS long, ST_Y(location::geometry) AS lat, open_time, close_time FROM technician WHERE username = $1", technician.Username)
	newTechnician := &config.Technician{}
	err := row.Scan(&newTechnician.TechnicianID, &newTechnician.Username, &newTechnician.ShopName, &newTechnician.MobileNumber, &newTechnician.Description, &newTechnician.Open, &newTechnician.Longitude, &newTechnician.Latitude, &newTechnician.OpenTime, &newTechnician.CloseTime)
	if err != nil {
		fmt.Println(err)
		return technician, fmt.Errorf("error fetching technician%v", err)
	}
	return *newTechnician, nil
}

func GetTechniciansLocation(db *sql.DB, request config.ServiceRequest) ([]config.AvailableTechnicians, error) {
	rows, err := db.Query(`
		SELECT username, shop_name, mobile_no, shop_desc, currently_open,
		ST_X(location::geometry) AS long, ST_Y(location::geometry) AS lat,
		ST_Distance(
			location, 
			ST_SetSRID(ST_MakePoint($1, $2), 4326)::geography
		) AS distance
		FROM technician
		WHERE ST_DWithin(
				location, 
				ST_SetSRID(ST_MakePoint($3, $4), 4326)::geography,
				$5
			)
		ORDER BY distance
		LIMIT $6;
 	`, request.Longitude, request.Latitude, request.Longitude, request.Latitude, request.Distance, request.NumberOfTechnicians)
	if err != nil {
		return nil, fmt.Errorf("error fetching technicians: %v", err)
	}
	defer rows.Close()

	var technicians []config.AvailableTechnicians

	for rows.Next() {
		var technician config.AvailableTechnicians
		err := rows.Scan(&technician.Username, &technician.ShopName, &technician.MobileNumber, &technician.Description, &technician.Open, &technician.Longitude, &technician.Latitude, &technician.Distance)
		if err != nil {
			return nil, fmt.Errorf("error scanning technician row: %v", err)
		}
		technicians = append(technicians, technician)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("error after iterating rows: %v", err)
	}

	return technicians, nil
}

func DeleteTechnician(db *sql.DB, technician config.Technician) (config.Technician, error) {
	_, err := db.Exec(
		"DELETE FROM technician WHERE username = $1", technician.Username,
	)
	if err != nil {
		return technician, fmt.Errorf("error deleting technician %v", err)
	}

	return technician, nil
}

func UpdateShopDetails(db *sql.DB, technician config.Technician) (config.Technician, error) {
	_, err := db.Exec(
		"UPDATE technician SET shop_name=$1, mobile_no=$2, shop_desc=$3, currently_open=$4, location=ST_SetSRID(ST_MakePoint($5, $6), 4326), open_time=$7, close_time=$8 WHERE username=$9",
		technician.ShopName, technician.MobileNumber, technician.Description, technician.Open, technician.Longitude, technician.Latitude, technician.OpenTime, technician.CloseTime, technician.Username,
	)
	if err != nil {
		fmt.Println(err)
		return technician, fmt.Errorf("error updating technician %v", err)
	}

	return technician, nil
}

func StoreRequest(db *sql.DB, request config.ServiceRequest, technicians []config.AvailableTechnicians) error {
	id := uuid.New().String()

	for i := 0; i < len(technicians); i++ {
		technician := technicians[i]
		_, err := db.Exec(
			"INSERT INTO temprequest (id, requestid, userid, technician) VALUES ($1, $2, $3, $4)",
			id, request.RequestID, request.UserID, technician.Username,
		)
		if err != nil {
			fmt.Println(err)
			return fmt.Errorf("error inserting request %v", err)
		}
	}

	return nil
}

func RemoveRequest(db *sql.DB, request config.ServiceRequest, flag bool) error {
	var err error
	if flag {
		_, err = db.Exec(
			"DELETE FROM temprequest WHERE requestid = $1", request.RequestID,
		)
	}else {
		_, err = db.Exec(
			"DELETE FROM temprequest WHERE requestid = $1 AND technician = $2", request.RequestID, request.Technician,
		)
	}
	if err != nil {
		return fmt.Errorf("error deleting request %v", err)
	}

	return nil
}
