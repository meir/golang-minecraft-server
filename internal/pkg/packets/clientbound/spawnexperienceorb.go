package clientbound

import "gomc/internal/pkg/packets"

type SpawnExperienceOrbPacket struct {
	*packets.Packet

	EntityID int `packets:"varint"`
	X float64 `packets:"double"`
	Y float64 `packets:"double"`
	Z float64 `packets:"double"`
	Count int16 `packets:"short"`
}

func NewSpawnExperienceOrbPacket(EntityID int, X, Y, Z float64, Count int16) *SpawnExperienceOrbPacket {
	return &SpawnExperienceOrbPacket{
		Packet:   &packets.Packet{
			Length:   0,
			PacketID: 0x01,
		},
		EntityID: EntityID,
		X:        X,
		Y:        Y,
		Z:        Z,
		Count:    Count,
	}
}
