package packet

type Packet struct {
	IncomingBuffer string
	StartBytes     string
	PacketLength   string
	ProtocolNumber string
	Content        string
	PacketSerial   string
	StopBytes      string
}

func (s Packet) SetIncomintBuffer() string {
	return ""
}
func (s Packet) SetStartBytes() string {
	return ""
}
func (s Packet) setPacketLength() string {
	return ""
}
func (s Packet) SetProtocolNumber() string {
	return ""
}

func (s Packet) SetContent() string {
	return ""
}

func (s Packet) SetPacketSerial() string {
	return ""
}
func (s Packet) SetStopBytes() string {
	return ""
}