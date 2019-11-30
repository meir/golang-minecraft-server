package clientbound

import "gomc/internal/pkg/packets"

type SpawnGlobalEntityPacket struct {
	*packets.Packet

	EntityID int `packets:"varint"`
	Type

}