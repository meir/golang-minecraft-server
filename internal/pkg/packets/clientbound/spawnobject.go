package clientbound

import "gomc/internal/pkg/packets"

type SpawnObjectPacket struct {
	*packets.Packet

	EntityID int `packets:"varint"`
	ObjectID packets.UUID `packets:"uuid"`
	Type int `packets:"varint"`

	X float64 `packets:"double"`
	Y float64 `packets:"double"`
	Z float64 `packets:"double"`

	Pitch int16 `packets:"short"`
	Yaw int16 `packets:"short"`

	Data int `packets:"int"`

	VelocityX int16 `packets:"short"`
	VelocityY int16 `packets:"short"`
	VelocityZ int16 `packets:"short"`
}

func NewSpawnObjectPacket(EntityID int, ObjectID packets.UUID, Type int, X, Y, Z float64, Pitch, Yaw int16, Data int, VelocityX, VelocityY, VelocityZ int16) *SpawnObjectPacket {
	return &SpawnObjectPacket{
		&packets.Packet{
			Length:   0,
			PacketID: 0x00,
		},
		EntityID,
		ObjectID,
		Type,
		X,
		Y,
		Z,
		Pitch,
		Yaw,
		Data,
		VelocityX,
		VelocityY,
		VelocityZ,
	}
}