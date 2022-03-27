package operations

import (
	"testing"
)

func TestOPTypeExaustiveAliases(t *testing.T) {
	got := len(OP_TYPE_ALIASES)
	want := int(OP_TYPE_COUNT)

	if got != want {
		t.Errorf("OP_TYPE_ALIASES has %d elements, want %d", got, want)
	}
}
