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
	"bool":  true,
	"dump":  true,
	"else":  true,
	"end":   true,
	"false": true,
	"float": true,
	"if":    true,
	"int":   true,
	"str":   true,
	"true":  true,
	"while": true,
}
