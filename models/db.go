package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
)

var Db *sql.DB // Note the sql package provides the namespace

func Init() {
	// var err error
	// Db, err = sql.Open("mysql", "root:winter2summer@tcp(127.0.0.1:3306)/vts_geo_tracker")
	// defer Db.Close()

	// if err != nil {
	// 	log.Fatal(err)
	// }
}
func dbConn() (db *sql.DB, e error) {

	//db, err := sql.Open("mysql", "root:winter2summer@tcp(127.0.0.1:3306)/vts_geo_tracker")

	//return db, err

	return sql.Open("mysql", "root:winter2summer@tcp(127.0.0.1:3306)/vehicletr")
}
