package handlers

import (
	ops "eslang/core/operations"
	"testing"
)

func TestExaustiveOperationsHandling(t *testing.T) {
	got := len(REGISTERED_OPERATIONS)
	want := int(ops.OP_TYPE_COUNT) - 1

	if got != want {
		t.Errorf("Expected %d operations, got %d", want, got)
	}
}
