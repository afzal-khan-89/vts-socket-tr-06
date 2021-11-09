package clientHandler

import (
	"encoding/hex"
	"fmt"
	"net"
	"strconv"

	"github.com/afzal-khan-89/vts-socket-tr-06.git/packet"
	"github.com/afzal-khan-89/vts-socket-tr-06.git/util"
)

func HandleLoginRequest(conn net.Conn, packet *packet.Packet, terminalId *string, dataStatus *string) error {
	*terminalId = packet.Content
	//_, err := conn.Write([]byte("dhoner matha ...\n"))
	responseForTerminal(conn, packet, dataStatus)
	return nil
}

func responseForTerminal(conn net.Conn, recvPacket *packet.Packet, dataType *string) error {
	fmt.Println("clientHandler : data type ", *dataType)
	fmt.Println("clientHandler : recvPacket.Protocol ", recvPacket.Protocol)
	if *dataType == "A" && (recvPacket.Protocol == "01" || recvPacket.Protocol == "13") {
		/* prepare response data */
		outgoingDataPacket := recvPacket.StartBytes                                  // initialize with start bits.
		responseDataLength := "05"                                                   //hex represent of decimal 5
		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, responseDataLength)      //push data length
		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, recvPacket.Protocol)     //push protocol no.
		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, recvPacket.PacketSerial) //push serial no

		/* generate and push error code */
		data_p := fmt.Sprint(responseDataLength, recvPacket.Protocol, recvPacket.PacketSerial)
		responseErrorHex, _ := hex.DecodeString(data_p)
		responseDataCRC := util.Checksum(responseErrorHex)                       //Error code in uint16
		outgoingDataErrorCode := strconv.FormatUint(uint64(responseDataCRC), 16) //Error code in string
		if len(outgoingDataErrorCode) == 3 {
			outgoingDataErrorCode = fmt.Sprint("0", outgoingDataErrorCode)
		}
		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, outgoingDataErrorCode)

		outgoingDataPacket = fmt.Sprint(outgoingDataPacket, recvPacket.StopBytes) //push stop bit
		/* send response to terminal */
		fmt.Println("OUTGOING :" + outgoingDataPacket)
		hexDataPacket, responseDataError := hex.DecodeString(outgoingDataPacket)
		/* set login status */

		fmt.Println("OUT-GOING PACKET : ", outgoingDataPacket)
		fmt.Println("OUT-GOING PACKET HED : ", hexDataPacket)

		if responseDataError == nil {
			_, writeError := conn.Write(hexDataPacket)
			if recvPacket.Protocol == "01" && writeError == nil {
				fmt.Println("Login success for: ")
				//loginState = true
			}
		} else {
			fmt.Println("Response Data Error for :")
			fmt.Println(responseDataError.Error())

		}
	}
	return nil
}
