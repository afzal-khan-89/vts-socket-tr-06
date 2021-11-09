package main

import (
	"encoding/hex"
	"fmt"
	"io"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

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
	var packet Packet
	var terminalId string
	var startBytes string

	for {
		err := conn.SetReadDeadline(time.Now().Add(5 * time.Second)) // set SetReadDeadline
		if err != nil {
			log.Println("SetReadDeadline failed:", err)
			return
		}

		buff := make([]byte, 8192)

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
			//dataLength := len(incomingDataArray)
			startBytes = incomingDataArray[i][:4]
			if startBytes == "7878" {
				packet := fmt.Sprint(incomingDataArray[i], "0d0a")
				fmt.Println("log: Packet  : ", packet)
				fmt.Println("log: packet length : ", len(packet))
				packetLength := packet[4:6]
				fmt.Println("log: Packet Length ", packetLength)
				incomintContent := packet[8 : len(packet)-12]
				fmt.Println("log: Incoming content  : ", incomintContent)
				packetCRC := packet[len(packet)-8 : len(packet)-4]
				fmt.Println("log: PacketCrc", packetCRC)

				stringToCheckCRC := packet[4 : len(incomintContent)+12]
				fmt.Println("log: crc string : ", stringToCheckCRC)
				fmt.Println("log: incomingPacketLength ", len(incomintContent))

				status, err := util.CRCcheck(stringToCheckCRC, packetCRC)
				if err != nil {
					fmt.Println("log: CRC error")
				}
				fmt.Println("log: status ", status)

				protocolNumber := packet[6:8]
				fmt.Println("log: Progocol :: ", protocolNumber)
				switch protocolNumber {
				case PROTOCOL_LOGIN_MESSAGE:
					fmt.Println("log: PROTOCOL_LOGIN_MESSAGE ")
					handleLoginRequest(conn, &incomintContent, &terminalId, &startBytes)
					fmt.Println("terminal Id : ", terminalId)
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
func handleLoginRequest(conn net.Conn, data *string, terminalId *string, startBytes *string) error {
	*terminalId = *data
	fmt.Println(fmt.Sprint("Terminal ID: ", *terminalId))
	_, err := conn.Write([]byte("dhoner matha ...\n"))
	return err
}
func handleHeartbitData(conn net.Conn) error {
	return nil
}
func responseForTerminal(startBytes *string, dataType *string, protocol *string) error {
	if *dataType == "A" && (*protocol == "01" || *protocol == "13") {
		/* prepare response data */
		var outGoingPacket string
		outGoingPacket = fmt.Sprint(outGoingPacket, startBytes)
		outGoingPacket = fmt.Sprint(outGoingPacket, "05")
		outGoingPacket = fmt.Sprint(outGoingPacket, protocol)
		outGoingPacket = fmt.Sprint(outGoingPacket, protocol)

		outgoingDataPacket := &startBytes                                         // initialize with start bits.
		responseDataLength := "05"                                                //hex represent of decimal 5
		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, responseDataLength)   //push data length
		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, incomingDataProtocol) //push protocol no.
		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, serialNo)             //push serial no
		/* generate and push error code */
		data_p := fmt.Sprint(responseDataLength, incomingDataProtocol, serialNo)
		responseErrorHex, _ := hex.DecodeString(data_p)
		responseDataCRC := Checksum(responseErrorHex, table)                     //Error code in uint16
		outgoingDataErrorCode := strconv.FormatUint(uint64(responseDataCRC), 16) //Error code in string
		if len(outgoingDataErrorCode) == 3 {
			outgoingDataErrorCode = fmt.Sprint("0", outgoingDataErrorCode)
		}
		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, outgoingDataErrorCode)

		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, stopBits) //push stop bit
		/* send response to terminal */
		fmt.Println("OUTGOING :" + outgoingDataPacket)
		hexDataPacket, responseDataError := hex.DecodeString(outgoingDataPacket)
		/* set login status */
		if responseDataError == nil {
			_, writeError := conn.Write(hexDataPacket)
			if incomingDataProtocol == "01" && writeError == nil {
				fmt.Println("Login success for: " + terminalId)
				loginState = true
			}
		} else {
			fmt.Println("Response Data Error for :" + terminalId)
			fmt.Println(responseDataError.Error())
			return
		}
	}
	return nil
}
