package core

import "testing"

func TestTypeExaustiveAliases(t *testing.T) {
	got := len(TYPE_ALIASES)
	want := int(TYPE_COUNT)

	if got != want {
		t.Errorf("TYPE_ALIASES has %d elements, want %d", got, want)
	}
}
