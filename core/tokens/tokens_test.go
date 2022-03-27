package tokens

import "testing"

func TestTokenTypeExaustiveAliases(t *testing.T) {
	got := len(TOKEN_ALIASES)
	want := int(TOKEN_TYPE_COUNT)

	if got != want {
		t.Errorf("TOKEN_ALIASES has %d entries, want %d", got, want)
	}
}
