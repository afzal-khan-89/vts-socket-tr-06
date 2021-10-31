package models

import "fmt"

type Vehicledata struct {
	Dtime     string
	Status    bool
	Imei      string
	Latitude  string
	Longitude string
}

func GetAllVehicleData() []Vehicledata {
	var vehicleData []Vehicledata
	db, err := dbConn()
	if err != nil {
		panic("fail to connect to database .")
	}
	defer db.Close()

	results, err := db.Query("SELECT * FROM gps_data")
	if err != nil {
		panic(err.Error()) // proper error handling instead of panic in your app
	}
	for results.Next() {
		var tag Vehicledata
		err = results.Scan(&tag.Dtime, &tag.Imei, &tag.Latitude, &tag.Longitude, &tag.Status)
		//err = results.Scan(&tag.ID, &tag.Imei, &tag.Dtime, &tag.Latitude, &tag.Longitude, &tag.Speed, &tag.Status)
		if err != nil {
			panic(err.Error()) // proper error handling instead of panic in your app
		}
		//fmt.Println(tag.Imei + "  " + tag.Longitude + "  " + tag.Latitude)
		vehicleData = append(vehicleData, tag)
		// and then print out the tag's Name attribute
		// log.Printf(tag.imei)
		// log.Printf(tag.latitude)
		// log.Printf(tag.longitude)
	}
	return vehicleData
}
func SaveGpsData(data Vehicledata) (e error) {
	db, err := dbConn()
	if err != nil {
		panic("fail to connect to database .")
	}
	defer db.Close()
	insert, err := db.Query("INSERT INTO Vehicles VALUES (Vehicledata.Dtime, Vehicledata.Status, Vehicledata.Imei, Vehicledata.Latitude, Vehicledata.Longitude")
	if err != nil {
		panic("fail to insert to database .")
	}
	fmt.Println("Insert success full ... ")
	insert.Close()
	return nil
}
