package main

import (
	"bufio"
	"fmt"
	"net"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	vehicle := Vehicledata{Dtime: nil, Status: 1, Imei: "234234234", Latitude: "23.423424343", Longitude: "90.2234324324"}

	fmt.Println("hello ....")
	ln, _ := net.Listen("tcp", "localhost:3030")
	fmt.Print("Waiting for clients ... ")

	for {
		conn, _ := ln.Accept()
		mess, _ := bufio.NewReader(conn).ReadString('\n')
		fmt.Println("message received : " + mess)
	}
}
