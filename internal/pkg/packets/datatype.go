package packets

type Chat struct {
	Text string 				`json:"text"`

	Bold string 				`json:"bold,omitempty"`
	Italic string 				`json:"italic,omitempty"`
	Underlined string 			`json:"underlined,omitempty"`
	Strikethrough string 		`json:"strikethrough,omitempty"`
	Obfuscated string 			`json:"obfuscated,omitempty"`
	Color string 				`json:"color,omitempty"`

	Insertion string 			`json:"insertion,omitempty"`

	ClickEvent struct {
		OpenURL string 			`json:"open_url,omitempty"`
		RunCommand string 		`json:"run_command,omitempty"`
		SuggestCommand string 	`json:"suggest_command,omitempty"`
		ChangePage string 		`json:"change_page,omitempty"`
	} 							`json:"clickEvent,omitempty"`

	HoverEvent struct {
		ShowText string 		`json:"show_text,omitempty"`
		ShowItem interface{} 	`json:"show_item,omitempty"` // NBT struct
		ShowEntity interface{} 	`json:"show_entity,omitempty"` // ^
	} 							`json:"hoverEvent,omitempty"`

	Extra []Chat 				`json:"extra,omitempty"`
}

type Identifier struct {
	Namespace string
	Thing string
}

type Position struct {
	X int32
	Y int32
	Z int32
}

const UUID_DEFAULT string = "00000000-0000-0000-0000-000000000000"

type UUID struct {
	Most uint64
	Least uint64

	Raw []byte

	Str string
}