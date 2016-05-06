package domain

import (
	"database/sql"
	"log"
)

//BusinessTags Model for storing the relationship between business id and tag id
type BusinessTags struct {
	Id int64 `json:"id"`
	TagId int64 `json:"tag_id"`
	BusinessId int64 `json:"business_id"`
}

type BusinessTagHTTPResponse struct {
	Err bool `json:"err"`
	Msg string `json:"msg"`
	BusinessTags BusinessTags `json:"business_tags"`
}


//AddBusinessTag Add business tag to the db
func AddBusinessTag(db *sql.DB, rel *BusinessTags) error {
	//todo improve error handling in case of non unique tag addition
	stmt, err := db.Prepare("INSERT INTO business_tags (id, tag_id, business_id) VALUES (?, ?, ?)")
	if err != nil {
		return err
	}

	res, err := stmt.Exec(rel.Id, rel.TagId, rel.BusinessId)
	if err != nil {
		return err
	}

	rel.Id, err = res.LastInsertId()
	return err
}

//BusinessTagDetails Details of all the tag of a businessId
type BusinessTagDetails struct {
	BusinessId int64 `json:"business_id"`
	Tags []string `json:"tags"`
}

//BusinessTagDetailsHTTPResponse HTTP response to get business tags
type BusinessTagDetailsHTTPResponse struct {
	Err bool `json:"err"`
	Msg string `json:"msg"`
	BusinessTags BusinessTagDetails `json:"business_tags"`
}


//GetTagsOfBusiness retrieve all tag names of a particular business
func GetTagsOfBusiness(db *sql.DB, businessId *int64) (*BusinessTagDetails, error)  {

	stmt, err := db.Prepare("SELECT tag.name FROM business_tags JOIN tag ON" +
		" business_tags.tag_id = tag.id WHERE business_id=?")
	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(*businessId)
	if err != nil {
		return nil, err
	}

	return parseRowsBusinessTag(rows, businessId)
}


func parseRowsBusinessTag(rows *sql.Rows, bID *int64) (*BusinessTagDetails, error) {
	var err error

	var bTag BusinessTagDetails
	bTag.BusinessId = *bID

	defer rows.Close()
	for rows.Next() {
		var tagName string
		err = rows.Scan(&tagName)
		if err != nil {
			log.Println("Could not scan row")
		} else {
			bTag.Tags = append(bTag.Tags, tagName)
		}
	}

	return &bTag, rows.Err()

}

