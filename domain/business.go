package domain

import (
	"database/sql"
	"log"
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
func AddBusiness(db *sql.DB, business *Business) (int64, error) {
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

type GetAllBusinessHTTPResponse struct {
	Err bool `json:"error"`
	Msg string `json:"msg"`
	Business []Business `json:"business"`
}

//GetAllBusiness Get all the business
func GetAllBusiness(db *sql.DB) ([]Business, error) {
	//todo check if using pointers will help here
	var businessArr []Business
	rows, err := db.Query("SELECT * FROM business")
	if err != nil {
		return businessArr, err
	}

	err = readAllBusinessRows(rows, &businessArr)
	return businessArr, err
}


func readAllBusinessRows(rows *sql.Rows,b *[]Business) error {
	var err error

	defer rows.Close()

	for rows.Next() {
		var business Business

		err = rows.Scan(&business.ID, &business.Name, &business.City, &business.Email,
			&business.Phone, &business.Latitude, &business.Longitude, &business.Rating)
		if err != nil {
			log.Println("Error Occured while Reading row")
		} else {
			*b = append(*b, business)
		}
	}

	rows.Close()
	return rows.Err()
}