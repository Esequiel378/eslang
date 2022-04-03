package core

type Type int

const (
	Bool Type = iota
	Float
	Int
	Nil
	String

	TYPE_COUNT
)

var TYPE_ALIASES = map[Type]string{
	Bool:   "bool",
	Float:  "float",
	Int:    "int",
	Nil:    "nil",
	String: "str",
}

func (t Type) String() string {
	return TYPE_ALIASES[t]
}

var RESERVED_WORDS = map[string]bool{
	"-rot":  true,
	"2drop": true,
	"2dup":  true,
	"2over": true,
	"2swap": true,
	"bool":  true,
	"do":    true,
	"drop":  true,
	"dump":  true,
	"dup":   true,
	"else":  true,
	"end":   true,
	"false": true,
	"float": true,
	"if":    true,
	"int":   true,
	"nip":   true,
	"over":  true,
	"rot":   true,
	"str":   true,
	"swap":  true,
	"true":  true,
	"tuck":  true,
	"while": true,
}
