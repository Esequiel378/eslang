package tokens

import (
	"testing"
)

func TestTokenExaustiveHandlers(t *testing.T) {
	got := len(REGISTERED_TOKENS)
	want := int(TOKEN_TYPE_COUNT)

	if got != want {
		t.Errorf("TOKEN_HANDLERS has %d entries, want %d", got, want)
	}
}
