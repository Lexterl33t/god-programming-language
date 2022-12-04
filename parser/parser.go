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

func NewParser(sourceCode string) Parser_t {
	return Parser_t{
		Lex: lexer.NewLexer(sourceCode),
	}
}

func Parse() {

}
