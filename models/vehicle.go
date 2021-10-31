package models

import (
	"fmt"
	"log"
	"time"
)

type RawData struct {
	DataTime  time.Time
	Status    bool
	Imei      string
	Latitude  string
	Longitude string
}

func InsertRawData(d RawData) error {
	db, err := dbConn()
	if err != nil {
		log.Println("insert device db connection e akam ghotse ...")
	}
	defer db.Close()
	qStmt, err := db.Prepare("INSERT INTO vehicles(imei, data_time, status, latitude, longitude) VALUES(?,?,?,?,?)")
	if err != nil {
		log.Println("insert device query banaite akam ghotse ...")
	}

	log.Println("-imei : " + d.Imei)
	log.Println("-latitude : " + d.Latitude)
	log.Println("-longitude : " + d.Longitude)

	res, err := qStmt.Exec(d.Imei, d.DataTime, d.Status, d.Latitude, d.Longitude)
	if err != nil {
		log.Println("device db te insert korte pare nai ... ")
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println("device insert kore last id receive korte pare nai ..")
	}
	fmt.Println("Insert id", id)
	// http.Redirect(w, r, "/", 301)
	return err
}
