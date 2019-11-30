package serverbound

import (
	"gomc/internal/pkg/packets"
)

type HandshakePacket struct {
	packets.Packet

	ProtocolVersion int	`packet:"varint"`
	ServerAddress string `packet:"string"`
	ServerPort uint16 `packet:"unsigned_short"`
	NextState int `packet:"varint"`
}

func NewHandshakePacket(src []byte) *HandshakePacket {
	packet := &HandshakePacket{
		Packet:          *packets.NewPacket(src),
	}
	packet.Parse(packet)
	return packet
}