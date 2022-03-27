package core

type Type int

const (
	Float Type = iota
	Int
	Nil
	String

	TYPE_COUNT
)

var TYPE_ALIASES = map[Type]string{
	Float:  "float",
	Int:    "int",
	Nil:    "nil",
	String: "str",
}

func (t Type) String() string {
	if alias, ok := TYPE_ALIASES[t]; ok {
		return alias
	}

	return "-unknown-"
}

// TODO: add tests to check the amount of RESERVED_WORDS
var RESERVED_WORDS = map[string]bool{
	"do":    true,
	"dump":  true,
	"dup":   true,
	"else":  true,
	"end":   true,
	"float": true,
	"if":    true,
	"int":   true,
	"str":   true,
	"while": true,
}
