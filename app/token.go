package main

import (
	"fmt"
)

type Token struct {
	token_type TokenType
	lexeme     string
	literal    any
	line       int
}

func stringifyToken(token Token) string {
	var literal string
	if token.literal == nil {
		literal = "null"
	}
	return fmt.Sprintf("%s %s %s", token.token_type, token.lexeme, literal)
}
