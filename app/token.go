package main

import (
	"fmt"
	"strconv"
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
	} else if value, ok := token.literal.(string); ok {
		literal = value
	} else if value, ok := token.literal.(int); ok {
		literal = strconv.Itoa(value)
	}
	return fmt.Sprintf("%s %s %s", token.token_type, token.lexeme, literal)
}
