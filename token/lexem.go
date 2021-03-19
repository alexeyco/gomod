package token

import "fmt"

type Lexem struct {
	Token Token
	Value string
}

func (l Lexem) String() string {
	return fmt.Sprintf("%s %s", l.Token, l.Value)
}
