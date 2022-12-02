package main

import (
	"fmt"
	"os"

	"github.com/lexterl33t/god-programming-language/lexer"
	"github.com/lexterl33t/god-programming-language/token"
)

func main() {

	sourceCode, _ := os.ReadFile("examples/test.god")

	lexer := lexer.NewLexer(string(sourceCode))

	token_i := lexer.GetToken()

	for token_i.Kind != token.EOF {
		fmt.Println(token_i)
		token_i = lexer.GetToken()
	}
}
