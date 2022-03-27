package core

import "testing"

func TestOPTypeExaustiveAliases(t *testing.T) {
	got := len(OP_TYPE_ALIASES)
	want := int(OP_TYPE_COUNT)

	if got != want {
		t.Errorf("OP_TYPE_ALIASES has %d elements, want %d", got, want)
	}
}

func TestTypeExaustiveAliases(t *testing.T) {
	got := len(TYPE_ALIASES)
	want := int(TYPE_COUNT)

	if got != want {
		t.Errorf("TYPE_ALIASES has %d elements, want %d", got, want)
	}
}
