package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"strings"
	"time"

	"github.com/afzal-khan-89/vts-socket-tr-06.git/packet"
	"github.com/afzal-khan-89/vts-socket-tr-06.git/util"
)

const (
	PROTOCOL_LOGIN_MESSAGE                          = "01"
	PROTOCOL_LOGIN_DATA                             = "12"
	PROTOCOL_HEARTBEAT_OR_STATUS_INFO               = "13"
	PROTOCOL_STRING_INFORMATION                     = "15"
	PROTOCOL_ALARM_DATA                             = "16"
	PROTOCOL_GPS_ADDRESS_QUERY_PACKET               = "1A"
	PROTOCOL_LBS_ADDRESS_QUERY_PACKET               = "17"
	PROTOCOL_DEMAND_INFO_SENT_TO_TERMINAL_BY_SERVER = "80"
	PROTOCOL_ONLINE_COMMAND_RESPONSE                = "15"
	PROTOCOL_TIME_CALIBRATION_PACKET                = "8A"
	PROTOCOL_INFORMATION_TRANSMISSION_PACKATE       = "94"
	PROTOCOL_COMMAND_FROM_SERVER_TO_TERMINAL        = "80"
)

func main() {
	fmt.Println("hello ....")
	ln, _ := net.Listen("tcp", "localhost:5050")
	fmt.Print("Waiting for clients ... ")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println("Error to accept client ... ")
		}
		//go handleClient(conn)
		go handleIt(conn)
	}

}
func handleIt(conn net.Conn) {
	var terminalId string

	for {
		packet := packet.Packet{}
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // set SetReadDeadline
		if err != nil {
			log.Println("SetReadDeadline failed:", err)
			return
		}

		buff := make([]byte, 1024)

		readLength, err := conn.Read(buff) // recv data
		if err != nil {
			//handle when client accidentally closed the connection
			if err == io.EOF {
				fmt.Println("--- connection dropped from client ---")
				return
			}
			continue
		}
		if readLength == 0 {
			fmt.Println("--Empty data ...")
			continue
		}
		fmt.Println(fmt.Sprint("DATA LENGTH: ", readLength))
		data := hex.EncodeToString(buff[:readLength])
		fmt.Println("DATA  :: ", data)
		//data sent from terminal split by stop bits
		incomingDataArray := strings.Split(data, "0d0a")
		fmt.Println(incomingDataArray)
		fmt.Println("log : incomingDataArrayLength : ", len(incomingDataArray))

		for i := 0; i < len(incomingDataArray)-1; i++ {
			fmt.Println("log: incomingDataArray ", incomingDataArray[i])
			packet.IncomingBuffer = fmt.Sprint(incomingDataArray[i], "0d0a")
			packet.StartBytes = incomingDataArray[i][:4]
			if packet.StartBytes == "7878" {
				fmt.Println("log: Incomint data  : ", packet.IncomingBuffer)
				fmt.Println("log: packet length : ", len(packet.IncomingBuffer))
				packet.PacketLength = packet.IncomingBuffer[4:6]
				fmt.Println("log: Packet Length ", packet.PacketLength)
				packet.Content = packet.IncomingBuffer[8 : len(packet.IncomingBuffer)-12]
				fmt.Println("log: Incoming content  : ", packet.Content)
				packet.CRC = packet.IncomingBuffer[len(packet.IncomingBuffer)-8 : len(packet.IncomingBuffer)-4]
				fmt.Println("log: PacketCrc", packet.CRC)

				stringToCheckCRC := packet.IncomingBuffer[4 : len(packet.Content)+12]
				fmt.Println("log: crc string : ", stringToCheckCRC)
				fmt.Println("log: incomingPacketLength ", len(packet.Content))

				status, err := util.CRCcheck(stringToCheckCRC, packet.CRC)
				if err != nil {
					fmt.Println("log: CRC error")
				}
				fmt.Println("log: status ", status)

				protocolNumber := packet.IncomingBuffer[6:8]
				fmt.Println("log: Progocol :: ", protocolNumber)
				switch protocolNumber {
				case PROTOCOL_LOGIN_MESSAGE:
					fmt.Println("log: PROTOCOL_LOGIN_MESSAGE ")
					handleLoginRequest(conn, &packet, &status, &terminalId)
					// fmt.Println("terminal Id : ", terminalId)
				case PROTOCOL_HEARTBEAT_OR_STATUS_INFO:
					fmt.Println("log: PROTOCOL_HEARTBEAT_OR_STATUS_INFO ")
				case PROTOCOL_ALARM_DATA:
					fmt.Println("log: PROTOCOL_ALARM_DATA ")
				default:
					fmt.Println("default")
				}
			}

		}
	}
}
func handleLoginRequest(conn net.Conn, packet *packet.Packet, status *string, tarminalId *string) error {
	*tarminalId = packet.Content
	fmt.Println(fmt.Sprint("Terminal ID: ", *tarminalId))
	_, err := conn.Write([]byte("dhoner matha ...\n"))
	return err
}
func handleHeartbitData(conn net.Conn) error {
	return nil
}
func responseForTerminal(startBytes *string, dataType *string, protocol *string) error {

	return nil
}
