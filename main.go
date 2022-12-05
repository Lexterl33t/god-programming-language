package main

import (
	"github.com/lexterl33t/god-programming-language/parser"
	"os"
)

func main() {

	sourceCode, _ := os.ReadFile("examples/test.god")

	parser_var := parser.NewParser()

	parser_var.Parse(string(sourceCode))
}
