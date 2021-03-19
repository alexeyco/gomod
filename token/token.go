package token

import "strconv"

const (
	Illegal = iota
	LineBreak
	Ident
	Comment // //

	literalBegin
	String // "value"
	Float  // 1.23
	literalEnd

	operatorBegin
	Implication      // =>
	LeftBracket      // [
	LeftParenthesis  // (
	RightBracket     // ]
	RightParenthesis // )
	operatorEnd

	keywordBegin
	Module
	Go
	Require
	Replace
	Retract
	keywordEnd
)

var tokens = [...]string{
	Illegal:   "illegal",
	LineBreak: "\n",
	Ident:     "ident",
	Comment:   "comment",

	Float: "float",

	Implication:      "=>",
	LeftBracket:      "[",
	LeftParenthesis:  "(",
	RightBracket:     "]",
	RightParenthesis: ")",

	Module:  "module",
	Go:      "go",
	Require: "require",
	Replace: "replace",
	Retract: "retract",
}

type Token uint

func (t Token) String() string {
	s := ""
	if 0 <= t && t < Token(len(tokens)) {
		s = tokens[t]
	}

	if s == "" {
		s = "token(" + strconv.Itoa(int(t)) + ")"
	}

	return s
}

func (t Token) IsLiteral() bool {
	return literalBegin < t && t < literalEnd
}

func (t Token) IsOperator() bool {
	return operatorBegin < t && t < operatorEnd
}

func (t Token) IsKeyword() bool {
	return keywordBegin < t && t < keywordEnd
}
