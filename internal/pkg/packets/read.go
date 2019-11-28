package packets

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"math"
	"strconv"
)

func(p *Packet) readByte() (interface{}, error) {
	b, err := p.buffer.ReadByte()
	if err != nil {
		return nil, err
	}
	return b, nil
}

func(p *Packet) readBytes(amount int) (interface{}, error) {
	var out []byte
	for i := 0; i < amount; i++ {
		data, err := p.readByte()
		if err != nil {
			return nil, err
		}
		out = append(out, data.(byte))
	}
	return out, nil
}

func(p *Packet) unreadByte() error {
	return p.buffer.UnreadByte()
}

func(p *Packet) readBool() (interface{}, error) {
	read, err := p.readByte()
	if err != nil {
		return false, err
	}else{
		return int(read.(byte)) == 1, nil
	}
}

func(p *Packet) readUnsignedByte() (interface{}, error) {
	return p.readByte()
}

func (p *Packet) readShort() (interface{}, error) {
	bytes, err := p.readBytes(2)
	if err != nil {
		return 0, err
	}
	var value int16
	value |= int16(bytes.([]byte)[0])
	value |= int16(bytes.([]byte)[1]) << 8
	return value, nil
}

func(p *Packet) readUnsignedShort() (interface{}, error) {
	bytes, err := p.readBytes(2)
	if err != nil {
		return 0, err
	}
	var value uint16 = binary.LittleEndian.Uint16(bytes.([]byte))
	return value, nil
}

func(p *Packet) readInt() (interface{}, error) {
	data, err := p.readByte()
	if err != nil {
		return 0, err
	}
	return int(data.(byte)), nil
}

func(p *Packet) readLong() (interface{}, error) {
	bytes, err := p.readBytes(8)
	if err != nil {
		return 0, err
	}
	var value int64
	value |= int64(bytes.([]byte)[0])
	value |= int64(bytes.([]byte)[1]) << 8
	value |= int64(bytes.([]byte)[2]) << 16
	value |= int64(bytes.([]byte)[3]) << 24
	value |= int64(bytes.([]byte)[4]) << 32
	value |= int64(bytes.([]byte)[5]) << 40
	value |= int64(bytes.([]byte)[6]) << 48
	value |= int64(bytes.([]byte)[7]) << 56
	return value, nil
}

func(p *Packet) readFloat() (interface{}, error) {
	bytes, err := p.readBytes(4)
	if err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint32(bytes.([]byte))
	return math.Float32frombits(bits), nil
}

func(p *Packet) readDouble() (interface{}, error) {
	bytes, err := p.readBytes(4)
	if err != nil {
		return 0, err
	}
	bits := binary.LittleEndian.Uint64(bytes.([]byte))
	return math.Float64frombits(bits), nil
}

func(p *Packet) readString() (interface{}, error) {
	strlen, err := p.readVarInt()
	if err != nil {
		return "", err
	}
	bytes, err := p.readBytes(strlen.(int))
	if err != nil {
		return "", err
	}
	return string(bytes.([]byte)), nil
}

func(p *Packet) readChat() (interface{}, error) { //needs Chat structure for return
	content, err := p.readString()
	if err != nil {
		return nil, err
	}
	if len(content.(string)) > 32767 {
		return nil, errors.New("Chat message reached maximum limit of " + strconv.Itoa(32767))
	}

	var output interface{} //needs Chat structure for return
	err = json.Unmarshal([]byte(content.(string)), &output)
	if err != nil {
		return nil, err
	}
	return output, nil
}

func(p *Packet) readIdentifier() (interface{}, error) { //requires identifier type
	content, err := p.readString()
	if err != nil {
		return "", err
	}
	if len(content.(string)) > 32767 {
		return "", errors.New("Chat message reached maximum limit of " + strconv.Itoa(32767))
	}

	return content, nil
}

func(p *Packet) readVarInt() (interface{}, error) {
	var numread int = 0
	var result int = 0

	for {
		read, err := p.readByte()
		if err != nil {
			return -1, err
		}
		var value int = int(read.(byte) & 0x7F)

		result |= value << (7 * numread)

		numread++
		if numread > 5 {
			return -1, errors.New("VarInt is too big")
		}
		if (read.(byte) & 0x80) == 0 {
			break
		}
	}
	return result, nil
}

func(p *Packet) readVarLong() (interface{}, error) {
	var numread int64 = 0
	var result int64 = 0

	for {
		read, err := p.readByte()
		if err != nil {
			return -1, err
		}
		var value int64 = int64(read.(byte) & 0x7F)
		result |= value << (7 * numread)

		numread++
		if numread > 10 {
			return -1, errors.New("VarLong is too big")
		}
		if (read.(byte) & 0x80) == 0 {
			break
		}
	}
	return result, nil
}