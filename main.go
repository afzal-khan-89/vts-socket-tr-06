package main

import (
	"bufio"
	"fmt"
	"net"
	"time"

	"github.com/afzal-khan-89/vts-socket-tr-06.git/models"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	currentTime := time.Now()
	vehicle := models.RawData{Imei: "234234234", DataTime: currentTime, Status: true, Latitude: "23.423424343", Longitude: "90.2234324324"}
	err := models.InsertRawData(vehicle)
	if err != nil {
		fmt.Println("error to same location data ...")
	}

	fmt.Println("hello ....")
	ln, _ := net.Listen("tcp", "localhost:3030")
	fmt.Print("Waiting for clients ... ")

	for {
		conn, _ := ln.Accept()
		mess, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("message received : " + mess)
	}
}
