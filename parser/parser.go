package parser

import "github.com/alexeyco/gomod/ast"

type Scanner interface {
}

type Parser struct {
	sc Scanner
}

func (p *Parser) Parse() (*ast.File, error) {
	return nil, nil
}

func New(sc Scanner) *Parser {
	return &Parser{
		sc: sc,
	}
}
