package scanner_test

import (
	"strings"
	"testing"

	"github.com/alexeyco/gomod/scanner"
)

func TestScanner_Next(t *testing.T) {
	t.Parallel()

	txt := `module foo.bar/fizz/buzz
go 1.2.3
`

	t.Run("", func(t *testing.T) {
		t.Parallel()

		sc := scanner.New(strings.NewReader(txt))

		t.Log(sc.Next())
		t.Log(sc.Next())
		t.Log(sc.Next())
		t.Log(sc.Next())
		t.Log(sc.Next())
	})
}
