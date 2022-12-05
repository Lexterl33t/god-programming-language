package parser

import (
	"fmt"
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

func (p *Parser_t) CheckToken(kind int) bool {
	if p.CurrentToken.Kind == kind {
		return true
	}

	return false
}

func (p *Parser_t) CheckPeekToken(kind int) bool {
	if p.PeekToken.Kind == kind {
		return true
	}

	return false
}

func (p *Parser_t) Match(kind int) {
	if p.CheckToken(kind) == false {
		panic(fmt.Sprintf("Expected %v, got %v", kind, p.CurrentToken.Kind))
	}

	p.NextToken()
}

func (p *Parser_t) NextToken() {

	p.CurrentToken = p.PeekToken

	p.PeekToken = p.Lex.GetToken()
}

func (p *Parser_t) Parse(sourceCode string) {

	p.Lex = lexer.NewLexer(sourceCode)

	p.NextToken()

	p.NextToken()

	p.State()
}

func (p *Parser_t) SkipNL() {

	if p.CheckToken(token.NEWLINE) == false {
		return
	}

	p.NextToken()

	p.SkipNL()
}

func (p *Parser_t) State() {

	if p.CheckToken(token.EOF) {
		return
	}

	p.SkipNL()

	fmt.Println(p.CurrentToken)

	p.NextToken()

	p.State()
}
