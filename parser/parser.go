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

func (p *Parser_t) CheckToken(tok token.Token_t) bool {
	if p.CurrentToken.Kind == tok.Kind {
		return true
	}

	return false
}

func (p *Parser_t) CheckPeekToken(tok token.Token_t) bool {
	if p.PeekToken.Kind == tok.Kind {
		return true
	}

	return false
}

func (p *Parser_t) Match() {

}

func (p *Parser_t) Parse(sourceCode string) {

	p.Lex = lexer.NewLexer(sourceCode)

}
