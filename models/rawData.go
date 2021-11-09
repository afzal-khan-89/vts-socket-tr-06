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
	Speed     float32
}

func CreateTableForDevice(deviceId string) error {
	db, err := dbConn()
	if err != nil {
		log.Println("insert device db connection e akam ghotse ...")
	}
	defer db.Close()
	device := "G_" + deviceId
	fmt.Println("device name : ", device)
	statement_1 := "CREATE TABLE IF NOT EXISTS " + device
	statement_2 := `(id INT NOT NULL AUTO_INCREMENT PRIMARY KEY UNIQUE, ` +
		`is_valid TINYINT(1) NOT NULL, ` +
		`data_time DATETIME NOT NULL,` +
		`latitude VARCHAR(14) NOT NULL, ` +
		`longitude VARCHAR(12) NOT NULL,` +
		`speed FLOAT NOT NULL,` +
		`fuel FLOAT NOT NULL, ` +
		`temperature FLOAT NOT NULL, ` +
		`ac TINYINT(1) NOT NULL,` +
		`vehicle_status INT NOT NULL)`

	fStatement := statement_1 + " " + statement_2

	fmt.Println("statememtn ---- :: ", fStatement)
	s, err := db.Exec(fStatement)

	//_, err := db.Exec("CREATE TABLE IF NOT EXISTS"+ device +"(id BIGINT NOT NULL AUTO_INCREMENT PRIMARY KEY, some_text TEXT NOT NULL)")
	if err != nil {
		panic(err)
	} else if s != nil {
		fmt.Println(`table ` + device + `Create successfully ... `)
	}
	return err
}
func InsertRawData(d RawData) error {
	db, err := dbConn()
	if err != nil {
		log.Println("insert device db connection e akam ghotse ...")
	}
	defer db.Close()
	//defer db.Close()

	if db == nil {
		fmt.Println("db is nil...")
	}

	qStmt, err := db.Prepare("INSERT INTO tbl99999999999999(imei, data_time, status, latitude, longitude, speed) VALUES(?,?,?,?,?,?)")
	if err != nil {
		log.Println("insert device query banaite akam ghotse ...")
	}

	log.Println("-data_time : ", d.DataTime)
	log.Println("-imei : ", d.Imei)
	log.Println("-latitude : ", d.Latitude)
	log.Println("-longitude : ", d.Longitude)
	log.Println("-status : ", d.Status)
	log.Println("-speed : ", d.Speed)

	fmt.Println("length of latitude :::::: ", len(d.Latitude))
	fmt.Println("length of longitude ::::: ", len(d.Longitude))

	//	_, ert := qStmt.Exec("3456346566", time.Now(), false, "23.23432434", "90.234433", 23.33)
	_, ert := qStmt.Exec(d.Imei, d.DataTime, d.Status, d.Latitude, d.Longitude, 12.23)

	if ert != nil {
		log.Println("device db te insert korte pare nai ... ")
	}
	// id, err := res.LastInsertId()
	// if err != nil {
	// 	log.Println("device insert kore last id receive korte pare nai ..")
	// }
	// fmt.Println("Insert id", id)
	return ert
}
