package main

import (
	"os"

	"github.com/smontalbano/parser/src/lexer"
)

func main() {
	bytes, _ := os.ReadFile("./examples/00.lang")
	tokens := lexer.Tokenize(string(bytes))

	for _, token := range tokens {
		token.Debug()
	}

}
