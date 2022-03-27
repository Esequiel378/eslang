package tokens

import (
	ops "eslang/core/operations"
	"testing"
)

func TestTokenTypeExaustiveAliases(t *testing.T) {
	got := len(TOKEN_ALIASES)
	want := int(TOKEN_TYPE_COUNT)

	if got != want {
		t.Errorf("TOKEN_ALIASES has %d entries, want %d", got, want)
	}
}

func TestTokenExaustiveHandlers(t *testing.T) {
	got := len(REGISTERED_TOKENS)
	want := int(TOKEN_TYPE_COUNT)

	if got != want {
		t.Errorf("TOKEN_HANDLERS has %d entries, want %d", got, want)
	}
}

func TestOPExaustiveHandlers(t *testing.T) {
	got := len(REGISTERED_TOKENS)
	want := int(ops.OP_TYPE_COUNT) - 1

	if got != want {
		t.Errorf("OP_HANDLERS has %d entries, want %d", got, want)
	}
}
