package scanner

import (
	"errors"
	"io"
	"text/scanner"

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
	panic("implement me")
}

func (s *Scanner) Peek() (*token.Lexem, error) {
	panic("implement me")
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
	s.Mode ^= scanner.SkipComments | scanner.ScanChars | scanner.ScanRawStrings

	// Don't skip line breaks
	s.Whitespace ^= 1 << '\n'

	return &Scanner{
		sc: s,
	}
}
