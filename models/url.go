package models

import (
	"fmt"
	"log"

	"github.com/ybuilds/slash/database"
	"github.com/ybuilds/slash/utils"
)

type Url struct {
	Id     int64  `json:"mapId"`
	Encode string `json:"encode"`
	Url    string `json:"url" validate:"required"`
}

var db = database.DB

func (u *Url) CreateMapping() (*Url, error) {
	query := `INSERT INTO url (url) VALUES (?)`
	res, err := db.Exec(query, u.Url)
	if err != nil {
		log.Println("error inserting url in database: ", err)
		return nil, err
	}

	u.Id, err = res.LastInsertId()
	u.Encode = utils.Base62Encoder(u.Id)
	if err != nil {
		fmt.Println("error fetching last inserted id of url: ", err)
		return nil, err
	}

	query = `INSERT INTO urlmap (encode, urlId) values (?, ?)`
	res, err = db.Exec(query, u.Encode, u.Id)
	if err != nil {
		log.Println("error creating mapping in database: ", err)
		return nil, err
	}

	return u, nil
}

func GetMapping(id int64) (*Url, error) {
	var url Url

	query := `SELECT um.mapId, um.encode, u.url FROM urlmap um INNER JOIN url u ON um.urlId = u.urlId WHERE u.urlId = ?`
	row := db.QueryRow(query, id)
	if row.Err() != nil {
		log.Println("error fetching url from database")
		return nil, row.Err()
	}

	err := row.Scan(&url.Id, &url.Encode, &url.Url)
	if err != nil {
		log.Println("error scanning cursor data: ", err)
		return nil, err
	}

	return &url, nil
}
