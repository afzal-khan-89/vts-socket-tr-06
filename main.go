package main

import (
	"fmt"

	"github.com/afzal-khan-89/vts-socket-tr-06.git/models"
)

func main() {
	//var rawData models.RawData
	models.CreateTableForDevice("1234545babu")
	fmt.Println("hello ....")
	/*ln, _ := net.Listen("tcp", "localhost:5050")
	fmt.Print("Waiting for clients ... ")
	conn, _ := ln.Accept()
	for {

		message, _ := bufio.NewReader(conn).ReadBytes('\n')
		data := string(message)
		fmt.Println("message received : " + string(data))
		dataArray := strings.Split(data, ",")
		fmt.Println("dataArray: ", dataArray)
		dataArray[1] = strings.TrimSuffix(dataArray[1], "\n")
		fmt.Println("lat: ", dataArray[0])
		fmt.Println("lon: ", dataArray[1])

		xt := reflect.TypeOf(dataArray).Kind()
		// yt := reflect.TypeOf(dataArray[1]).Kind()
		fmt.Printf("%T: %s\n", xt, xt)
		// fmt.Printf("%T: %s\n", yt, yt)

		rawData = models.RawData{Imei: "25878945612325", DataTime: time.Now(), Status: true, Latitude: dataArray[0], Longitude: dataArray[1], Speed: 12.22}
		err := models.InsertRawData(rawData)
		if err != nil {
			fmt.Println("error to same location data ...")
		}
	}
	*/
}
