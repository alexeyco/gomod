package scanner

import (
	"errors"
	"io"
	"text/scanner"
	"unicode"

	"github.com/alexeyco/gomod/token"
)

var (
	ErrEOF        = errors.New("EOF")
	ErrUnexpected = errors.New("unexpected")
)

type Scanner struct {
	sc scanner.Scanner
}

func (s *Scanner) Pos() token.Position {
	return token.Position{
		Line:   s.sc.Line,
		Column: s.sc.Column,
	}
}

func (s *Scanner) Next() (*token.Lexem, error) {
	return s.lex(s.sc.Scan(), s.sc.TokenText())
}

func (s *Scanner) Peek() (*token.Lexem, error) {
	return s.lex(s.sc.Peek(), s.sc.TokenText())
}

func (s *Scanner) lex(ch rune, text string) (lx *token.Lexem, err error) {
	l := token.Lexem{
		Text: text,
	}

	switch ch {
	case scanner.Ident:
		switch text {
		case "module":
			l.Token = token.Module
		case "go":
			l.Token = token.Go
		case "require":
			l.Token = token.Require
		case "replace":
			l.Token = token.Replace
		case "retract":
			l.Token = token.Retract
		default:
			l.Token = token.Ident
		}
	case '\n':
		l.Token = token.LineBreak
	case scanner.Comment:
		l.Token = token.Comment
	case scanner.Float:
		l.Token = token.Float
	case ',':
		l.Token = token.Comma
	case '[':
		l.Token = token.LeftBracket
	case '(':
		l.Token = token.LeftParenthesis
	case ']':
		l.Token = token.RightBracket
	case ')':
		l.Token = token.RightParenthesis
	case scanner.EOF:
		err = ErrEOF
	}

	if err == nil {
		lx = &l
	}

	return
}

func (s *Scanner) Skip(t token.Token, n int) (*token.Lexem, error) {
	panic("implement me")
}

func (s *Scanner) Want(t token.Token, n int) (*token.Lexem, error) {
	panic("implement me")
}

func New(r io.Reader) *Scanner {
	var s scanner.Scanner

	s.Init(r)

	// Don't skip commentaries but skip raw strings and chars
	s.Mode ^= scanner.SkipComments | scanner.ScanChars | scanner.ScanRawStrings | scanner.ScanFloats

	// Don't skip line breaks
	s.Whitespace ^= 1 << '\n'

	s.IsIdentRune = func(ch rune, i int) bool {
		return unicode.IsLetter(ch) || unicode.IsDigit(ch) || ((ch == '/' || ch == '.') && i > 0)
	}

	return &Scanner{
		sc: s,
	}
}
