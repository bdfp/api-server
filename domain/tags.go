package domain
import (
	"database/sql"
	"log"
)

//Tag Model for tag
type Tag struct {
	Id int64 `json:"id"`
	Name string `json:"name"`
}

//AddTag add a new tag
func AddTag (db *sql.DB, tag Tag) (Tag, error) {
	stmt, err := db.Prepare("INSERT INTO tag (name) VALUES (?)")
	if err != nil {
		return tag, err
	}

	res, err := stmt.Exec(tag.Name)
	if err != nil {
		return tag, err
	}

	tag.Id, err = res.LastInsertId()
	return tag, err
}

//todo create a generic method to get all rows of a database

//getAllTags get all tags
func GetAllTags (db *sql.DB) ([]Tag, error) {
	var tags []Tag

	rows, err := db.Query("SELECT * FROM tag")
	if err != nil {
		return tags, err
	}

	err = readAllTagRows(rows, &tags)
	return tags, err
}

func readAllTagRows(rows *sql.Rows, t *[]Tag) error {
	var err error

	defer rows.Close()

	for rows.Next() {
		var tag Tag

		if err = rows.Scan(&tag.Id, &tag.Name); err != nil {
			log.Println("Could not scan from row", t)
		} else {
			*t = append(*t, tag)
		}
	}

	return rows.Err()
}