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

var RESERVED_WORDS = map[string]bool{
	"dump":  true,
	"else":  true,
	"end":   true,
	"float": true,
	"if":    true,
	"int":   true,
	"str":   true,
}
