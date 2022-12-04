package lexer

import (
	"log"
	"unicode"

	"fmt"

	"github.com/lexterl33t/god-programming-language/token"
	"github.com/lexterl33t/god-programming-language/utils"
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

	if lex.CurrentChar == '+' {

		if lex.Peek() == '+' {

			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)),
				token.INCR,
			)

		} else {

			token_i = token.NewToken(string(lex.CurrentChar), token.PLUS)

		}

	} else if lex.CurrentChar == '-' {

		if lex.Peek() == '>' {

			lastchar := lex.CurrentChar

			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)),
				token.ARROWINITFUNC,
			)

		} else {

			token_i = token.NewToken(string(lex.CurrentChar), token.MINUS)

		}

	} else if lex.CurrentChar == '[' {

		token_i = token.NewToken(string(lex.CurrentChar), token.HOOKBEG)

	} else if lex.CurrentChar == ']' {

		token_i = token.NewToken(string(lex.CurrentChar), token.HOOKEND)

	} else if lex.CurrentChar == '{' {

		token_i = token.NewToken(string(lex.CurrentChar), token.BRACEBEG)

	} else if lex.CurrentChar == '}' {

		token_i = token.NewToken(string(lex.CurrentChar), token.BRACEEND)

	} else if lex.CurrentChar == '*' {

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

	} else if lex.CurrentChar == '=' {

		if lex.Peek() == '=' {

			lastChar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastChar), string(lex.CurrentChar)),
				token.DOUBLEEQ,
			)

		} else {

			token_i = token.NewToken(string(lex.CurrentChar), token.EQ)
		}

	} else if lex.CurrentChar == '!' {

		if lex.Peek() == '=' {

			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)),
				token.NOTEQ,
			)

		}

	} else if lex.CurrentChar == '>' {

		if lex.Peek() == '=' {

			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)),
				token.ABOVEEQ,
			)

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

	} else if lex.CurrentChar == '<' {
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

	} else if lex.CurrentChar == '&' {

		if lex.Peek() == '&' {

			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)), token.COMPAND)

		} else {

			token_i = token.NewToken(string(lex.CurrentChar), token.LOGICALAND)

		}

	} else if lex.CurrentChar == '|' {

		if lex.Peek() == '|' {

			lastchar := lex.CurrentChar
			lex.NextChar()

			token_i = token.NewToken(
				fmt.Sprintf("%v%v", string(lastchar), string(lex.CurrentChar)),
				token.COMPOR,
			)

		} else {

			token_i = token.NewToken(string(lex.CurrentChar), token.LOGICALOR)

		}

	} else if unicode.IsDigit(rune(lex.CurrentChar)) {

		startpos := lex.Position

		for unicode.IsDigit(rune(lex.CurrentChar)) || lex.CurrentChar == '.' {
			lex.NextChar()
		}

		var number string = lex.Source[startpos:lex.Position]

		token_i = token.NewToken(number, token.NUMBER)

	} else if lex.CurrentChar == '"' {

		lex.NextChar()

		startPos := lex.Position

		for lex.CurrentChar != '"' {
			lex.NextChar()
		}

		var str string = lex.Source[startPos:lex.Position]

		token_i = token.NewToken(str, token.STRING)

	} else if lex.CurrentChar == '(' {

		token_i = token.NewToken(string(lex.CurrentChar), token.PARENT)

	} else if lex.CurrentChar == ')' {

		token_i = token.NewToken(string(lex.CurrentChar), token.PARENT)

	} else if lex.CurrentChar == ',' {

		token_i = token.NewToken(string(lex.CurrentChar), token.SEPARATOR)

	} else if utils.IsAlpha(string(lex.CurrentChar)) {

		startPos := lex.Position

		for utils.IsAlpha(string(lex.CurrentChar)) {
			lex.NextChar()
		}

		var keyword_found string = lex.Source[startPos:lex.Position]

		if CheckIfKeyWord(keyword_found) {

			token_i = token.NewToken(keyword_found, token.Keyword[keyword_found])

		} else {

			token_i = token.NewToken(keyword_found, token.IDENT)

		}
	} else if lex.CurrentChar == '\n' {

		token_i = token.NewToken(string(lex.CurrentChar), token.NEWLINE)

	} else if lex.CurrentChar == 00 {

		token_i = token.NewToken(string(lex.CurrentChar), token.EOF)

	} else {

		log.Fatalln(fmt.Sprintf("%v unknow token", lex.CurrentChar))

	}

	lex.NextChar()

	return token_i
}

func CheckIfKeyWord(keyword string) bool {
	if _, ok := token.Keyword[keyword]; ok {
		return true
	}

	return false
}
