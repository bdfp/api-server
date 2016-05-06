package domain

import (
	"database/sql"
)

//Business Model for the db table
type Business struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	City      string `json:"city"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Latitude  string `json:"latitude"`
	Longitude string `json:"longitude"`
	Rating    string `json:"rating"`
}

//BusinessHTTPResponse Successful API Response
type BusinessHTTPResponse struct {
	Err      bool     `json:"err"`
	Msg      string   `json:"msg"`
	Business Business `json:"business"`
}

//AddBusiness Add the provided business to DB
func AddBusiness(db *sql.DB, business Business) (int64, error) {
	stmt, err := db.Prepare("INSERT INTO business (name, city, primary_email, primary_phone," +
		" latitude, longitude, overall_rating) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		return -1, err
	}

	res, err := stmt.Exec(business.Name, business.City, business.Email, business.Phone,
		business.Latitude, business.Longitude, business.Rating)
	if err != nil {
		return -1, err
	}

	business.ID, err = res.LastInsertId()
	if err != nil {
		return -1, err
	}

	return business.ID, nil
}