package token_test

import (
	"testing"

	"github.com/alexeyco/gomod/token"
)

func TestPosition_String(t *testing.T) {
	t.Parallel()

	pos := token.Position{
		Line:   123,
		Column: 456,
	}

	expected := "123:456"
	given := pos.String()

	if expected != given {
		t.Errorf(`Should be "%s", "%s" given`, expected, given)
	}
}
