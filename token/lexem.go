package token

import "fmt"

type Lexem struct {
	Token Token
	Text  string
}

func (l Lexem) String() string {
	return fmt.Sprintf("%s %s", l.Token, l.Text)
}
