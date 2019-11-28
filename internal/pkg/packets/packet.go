package packets

import (
	"bytes"
	"fmt"
	"go/types"
	"os"
	"reflect"
	"strconv"
	"text/tabwriter"
)

type Packet struct {
	Length int		`packet:"varint"`
	PacketID int	`packet:"varint"`

	source []byte
	buffer *bytes.Buffer
}

type PacketType struct {
	Type types.BasicKind
	Length int
}

func NewPacket(source []byte) *Packet {
	return &Packet{
		source: source,
		buffer: bytes.NewBuffer(source),
	}
}

func Parse(source []byte) *Packet {
	var packet = NewPacket(source)
	packet.Parse(packet)
	return packet
}

func(p *Packet) Print(in interface{}) {
	var value = reflect.TypeOf(in)
	header := fmt.Sprint("--------------- Printing significant values of ", value.Name(), " ---------------")
	println(header)
	w := tabwriter.NewWriter(os.Stdout, len(header)/4, 0, 4, ' ', tabwriter.TabIndent)
	for k, v := range p.GetValues(in) {
		var field = k
		var tag = v.Type
		if tag != "" {
			f := p.getFunction(tag)
			fieldValue := reflect.ValueOf(v.Value)
			val := fieldValue.String()
			switch fieldValue.Kind() {
			case reflect.Int:
				val = fmt.Sprintf("%d", int(fieldValue.Int()))
				break
			case reflect.Int8:
				val = fmt.Sprintf("%d", int8(fieldValue.Int()))
				break
			case reflect.Int16:
				val = fmt.Sprintf("%d", int16(fieldValue.Int()))
				break
			case reflect.Int32:
				val = fmt.Sprintf("%d", int32(fieldValue.Int()))
				break
			case reflect.Int64:
				val = fmt.Sprintf("%d", fieldValue.Int())
				break
			case reflect.Uint:
				val = fmt.Sprintf("%d", uint(fieldValue.Uint()))
				break
			case reflect.Uint8:
				val = fmt.Sprintf("%d", uint8(fieldValue.Uint()))
				break
			case reflect.Uint16:
				val = fmt.Sprintf("%d", uint16(fieldValue.Uint()))
				break
			case reflect.Uint32:
				val = fmt.Sprintf("%d", uint32(fieldValue.Uint()))
				break
			case reflect.Uint64:
				val = fmt.Sprintf("%d", fieldValue.Uint())
				break
			case reflect.Float32:
				val = fmt.Sprintf("%f", fieldValue.Float())
				break
			case reflect.Float64:
				val = fmt.Sprintf("%f", fieldValue.Float())
				break
			case reflect.Bool:
				val = strconv.FormatBool(fieldValue.Bool())
				break
			default:
				val = fieldValue.String()
			}
			if f != nil {
				_, err := fmt.Fprint(w, fmt.Sprint(field, "\t", tag, "\t", fieldValue.Type().String(), "\t", val, "\n"))
				if err != nil {
					continue
				}
			}
		}
	}
	err := w.Flush()
	if err != nil {
		println("Error while trying to print packet...")
		println(err.Error())
	}
}