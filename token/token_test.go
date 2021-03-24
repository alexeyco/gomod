package token_test

import (
	"testing"

	"github.com/alexeyco/gomod/token"
)

func TestToken_String(t *testing.T) {
	t.Parallel()

	testData := [...]struct {
		token    token.Token
		expected string
	}{
		{token: token.Illegal, expected: "illegal"},
		{token: token.LineBreak, expected: "\n"},
		{token: token.Ident, expected: "ident"},
		{token: token.Comment, expected: "comment"},

		{token: token.Float, expected: "float"},

		{token: token.Implication, expected: "=>"},
		{token: token.Comma, expected: ","},
		{token: token.LeftBracket, expected: "["},
		{token: token.LeftParenthesis, expected: "("},
		{token: token.RightBracket, expected: "]"},
		{token: token.RightParenthesis, expected: ")"},

		{token: token.Module, expected: "module"},
		{token: token.Go, expected: "go"},
		{token: token.Require, expected: "require"},
		{token: token.Replace, expected: "replace"},
		{token: token.Retract, expected: "retract"},
	}

	for _, testDatum := range testData {
		testDatum := testDatum

		t.Run(testDatum.token.String(), func(t *testing.T) {
			t.Parallel()

			given := testDatum.token.String()
			if testDatum.expected != given {
				t.Errorf(`Should be "%s", "%s" given`, testDatum.expected, given)
			}
		})
	}
}
