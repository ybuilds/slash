package models

import (
	"log"

	"github.com/ybuilds/slash/database"
)

type Url struct {
	Id     int64  `json:"mapId"`
	Encode string `json:"encode"`
	Url    string `json:"url" validate:"required"`
}

var db = database.DB

func (u *Url) CreateMapping() (int64, error) {
	query := `INSERT INTO url (encode, url) VALUES (?, ?)`

	res, err := db.Exec(query, u.Encode, u.Url)
	if err != nil {
		log.Println("error creating mapping in database: ", err)
		return -1, err
	}

	return res.LastInsertId()
}

func GetMapping(id int) (*Url, error) {
	query := `SELECT encode, url FROM url WHERE id = ?`

	var url Url

	row := db.QueryRow(query, id)
	if row.Err() != nil {
		log.Println("error fetching mapping from database")
		return nil, row.Err()
	}

	err := row.Scan(&url.Encode, &url.Url)
	if err != nil {
		log.Println("error scanning cursor data: ", err)
		return nil, err
	}

	return &url, nil
}
