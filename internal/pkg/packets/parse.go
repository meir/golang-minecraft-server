package packets

import (
	"reflect"
)

type ReadingFunc func () (interface{}, error)

func(p *Packet) getFunction(in string) *ReadingFunc {
	var funcs = map[string]ReadingFunc{
		"byte": 			ReadingFunc(p.readByte),
		"bool": 			ReadingFunc(p.readBool),
		"unsigned_byte": 	ReadingFunc(p.readUnsignedByte),
		"short": 			ReadingFunc(p.readShort),
		"unsigned_short": 	ReadingFunc(p.readUnsignedShort),
		"int": 				ReadingFunc(p.readInt),
		"long": 			ReadingFunc(p.readLong),
		"float": 			ReadingFunc(p.readFloat),
		"double": 			ReadingFunc(p.readDouble),
		"string": 			ReadingFunc(p.readString),
		"chat": 			ReadingFunc(p.readChat),
		"identifier": 		ReadingFunc(p.readIdentifier),
		"varint": 			ReadingFunc(p.readVarInt),
		"varlong": 			ReadingFunc(p.readVarLong),
	}
	if f, ok := funcs[in]; ok {
		return &f
	}else{
		return nil
	}
}

type ParserStruct struct {
	Type string
	Value interface{}
}

func(p *Packet) Parse(in interface{}) {
	values := p.GetValues(in)

	for k, v := range values {
		f := p.getFunction(v.Type)
		val, err := (*f)()
		if err != nil {
			println(err.Error())
			continue
		}
		values[k].Value = val
	}

	p.SetValues(in, values)
}

func(p *Packet) GetValues(in interface{}) map[string]*ParserStruct {

	if reflect.ValueOf(in).Type().Kind() != reflect.Ptr {
		println("(*Packet).GetValues(interface) map[string]*ParserStruct expected pointer to packet!")
		return map[string]*ParserStruct{}
	}

	values := p.getValues(p)
	packetValues := p.getValues(in)
	for k, v := range packetValues {
		values[k] = v
	}

	//fmt.Println("map:", packetValues)

	return values
}

func(p *Packet) getValues(in interface{}) map[string]*ParserStruct {
	var values = map[string]*ParserStruct{}

	if reflect.ValueOf(in).Type().Kind() != reflect.Ptr {
		println("(*Packet).getValues(interface) map[string]*ParserStruct expected pointer to packet!")
		return values
	}

	typeValue := reflect.ValueOf(in)
	v := reflect.ValueOf(typeValue.Interface())
	indirect := reflect.Indirect(v)
	structure := indirect.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if tag, ok := field.Tag.Lookup("packet"); ok {
			if f := p.getFunction(tag); f != nil {
				values[field.Name] = &ParserStruct{
					Type:  tag,
					Value: v.Elem().Field(i).Interface(),
				}
			}
		}
	}
	return values
}

func(p *Packet) SetValues(in interface{}, values map[string]*ParserStruct) {

	if reflect.ValueOf(in).Type().Kind() != reflect.Ptr {
		println("(*Packet).SetValues(interface, map[string]*ParserStruct) expected pointer to packet!")
		return
	}

	p.setValues(p, values)
	p.setValues(in, values)
}

func(p *Packet) setValues(in interface{}, values map[string]*ParserStruct) {
	if reflect.ValueOf(in).Type().Kind() != reflect.Ptr {
		println("(*Packet).setValues(interface, map[string]*ParserStruct) expected pointer to packet!")
		return
	}

	typeValue := reflect.ValueOf(in)
	v := reflect.ValueOf(typeValue.Interface())
	indirect := reflect.Indirect(v)
	structure := indirect.Type()
	for i := 0; i < structure.NumField(); i++ {
		field := structure.Field(i)
		if tag, ok := field.Tag.Lookup("packet"); ok {
			if f := p.getFunction(tag); f != nil {
				if val := values[field.Name]; val != nil {
					if v.Elem().Field(i).IsValid() && v.Elem().Field(i).CanSet() {
						v.Elem().Field(i).Set(reflect.ValueOf(val.Value))
					}
				}
			}
		}
	}
}