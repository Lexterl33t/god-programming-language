package lexer

import (
	"log"

	"fmt"

	"github.com/lexterl33t/god-programming-language/token"
)

type Lexer_t struct {
	CurrentChar byte
	Position    int
	End         int
	Source      string
}

func NewLexer(sourceCode string) *Lexer_t {
	return &Lexer_t{
		CurrentChar: ' ',
		Position:    -1,
		End:         0,
		Source:      sourceCode + "\n",
	}
}

func (lex *Lexer_t) NextChar() {
	lex.Position++
	if lex.Position >= len(lex.Source) {
		lex.CurrentChar = 00
	} else {
		lex.CurrentChar = lex.Source[lex.Position]
	}
}

func (lex *Lexer_t) Peek() byte {
	if lex.Position+1 >= len(lex.Source) {
		return 00
	}

	return lex.Source[lex.Position+1]
}

func (lex *Lexer_t) SkipWhiteSpace() {
	if lex.CurrentChar != ' ' && lex.CurrentChar != '\t' && lex.CurrentChar != '\r' {
		return
	}

	lex.NextChar()
	lex.SkipWhiteSpace()
}

func (lex *Lexer_t) SkipComments() {

	if lex.CurrentChar == '#' {
		for lex.CurrentChar != '\n' {
			lex.NextChar()
		}
	}

}

func (lex *Lexer_t) GetToken() token.Token_t {
	var token_i token.Token_t

	lex.SkipWhiteSpace()
	lex.SkipComments()

	switch lex.CurrentChar {
	case '+':
		token_i = token.NewToken(string(lex.CurrentChar), token.PLUS)
		break
	case '-':
		token_i = token.NewToken(string(lex.CurrentChar), token.MINUS)
		break
	case '*':

		if lex.Peek() == '*' {
			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)),
				token.POWEROP,
			)
		} else {
			token_i = token.NewToken(string(lex.CurrentChar), token.ASTERISQ)
		}
		break
	case '=':
		if lex.Peek() == '=' {
			lastChar := lex.CurrentChar
			lex.NextChar()
			token_i = token.NewToken(fmt.Sprintf("%v%v", string(lastChar), string(lex.CurrentChar)), token.DOUBLEEQ)
		} else {
			token_i = token.NewToken(string(lex.CurrentChar), token.EQ)
		}
		break
	case '!':
		if lex.Peek() == '=' {
			lastchar := lex.CurrentChar
			lex.NextChar()
			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)), token.NOTEQ)
		}
		break
	case '>':
		if lex.Peek() == '=' {
			lastchar := lex.CurrentChar
			lex.NextChar()
			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)), token.ABOVEEQ)
		} else if lex.Peek() == '>' {
			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)),
				token.LOGICALRIGHTBITSHIFT,
			)
		} else {
			token_i = token.NewToken(string(lex.CurrentChar), token.ABOVE)
		}
		break
	case '<':
		if lex.Peek() == '=' {
			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)), token.LESSEQ)
		} else if lex.Peek() == '<' {
			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)),
				token.LOGICALLEFTBITSHIFT,
			)
		} else {
			token_i = token.NewToken(string(lex.CurrentChar), token.LESS)
		}
		break
	case '&':
		if lex.Peek() == '&' {
			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)), token.COMPAND)
		} else {
			token_i = token.NewToken(string(lex.CurrentChar), token.LOGICALAND)
		}
		break
	case '|':
		if lex.Peek() == '|' {
			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)), token.COMPOR)
		} else {
			token_i = token.NewToken(string(lex.CurrentChar), token.LOGICALOR)
		}
		break
	case '\n':
		token_i = token.NewToken(string(lex.CurrentChar), token.NEWLINE)
		break
	case 00:
		token_i = token.NewToken(string(lex.CurrentChar), token.EOF)
		break
	default:
		log.Fatalln(fmt.Sprintf("%v UNknow token", lex.CurrentChar))
	}

	lex.NextChar()

	return token_i
}
