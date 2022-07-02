package src

import (
	"errors"
	"fmt"
)

type Program []Object

func Parse(tokens []Token) Program {
	if len(tokens) == 0 {
		return nil
	}
	var program Program
	if !tokens[0].compare(Token{
		Type:    LPAREN_T,
		Content: "(",
		x:       0,
		y:       0,
	}) {
		parserErr(tokens[0], errors.New("should be ("))
	}
	return program
}

func parserErr(token Token, err error) {
	addErr(errors.New(fmt.Sprintf("> parser error: %v\n\t> content = %s, x = %d, y = %d\n", err, token.Content, token.x, token.y)))
}
