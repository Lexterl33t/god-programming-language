package parser

import (
	"github.com/lexterl33t/god-programming-language/lexer"
	"github.com/lexterl33t/god-programming-language/token"
)

type Parser_t struct {
	Lex *lexer.Lexer_t

	CurrentToken token.Token_t
	PeekToken    token.Token_t
}

func NewParser() *Parser_t {
	return &Parser_t{}
}

func (p *Parser_t) Parse(sourceCode string) {
	p.Lex = lexer.NewLexer(sourceCode)
}
